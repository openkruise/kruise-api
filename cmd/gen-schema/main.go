package main

import (
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	semver "github.com/blang/semver/v4"
	"k8s.io/kube-openapi/pkg/common"
	kubeopenapiutil "k8s.io/kube-openapi/pkg/util"
	"k8s.io/kube-openapi/pkg/validation/spec"

	"github.com/openkruise/kruise-api/pkg/apis"
	"github.com/openkruise/kruise-api/pkg/kruise"
)

var K8SVersions = []int{18, 21, 24}

type xKubernetesGroupVersionKind struct {
	Group   string `json:"group"`
	Kind    string `json:"kind"`
	Version string `json:"version"`
}

type gvkMeta struct {
	XKubernetesGroupVersionKind []xKubernetesGroupVersionKind `json:"x-kubernetes-group-version-kind"`
}

type k8sGvkMapping struct {
	Definitions map[string]gvkMeta `json:"definitions"`
}

type openAPISchema struct {
	spec.Schema
	XKubernetesGroupVersionKind []xKubernetesGroupVersionKind `json:"x-kubernetes-group-version-kind"`
}

func (s openAPISchema) MarshalJSON() ([]byte, error) {
	b1, err := json.Marshal(s.Schema)
	if err != nil {
		return nil, fmt.Errorf("schema %v", err)
	}

	if s.XKubernetesGroupVersionKind != nil {
		b2, err := json.Marshal(s.XKubernetesGroupVersionKind)
		if err != nil {
			return nil, fmt.Errorf("x-kubernetes-group-version-kind %v", err)
		}
		b1 = append(b1[:len(b1)-1], fmt.Sprintf(",\"x-kubernetes-group-version-kind\":%s}", string(b2))...)
	}
	return b1, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func generateOpenApiSchema(outputPath string, GetOpenAPIDefinitions func(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition) error {
	// We replace the generated names with group specific names aka kruise is `apps.kruise.io` instead of the real
	// group kind because within all the openkruise projects we have overlapping types due to all openkruise projects being under the same
	// kruise.io group. Kustomize does not care about the name as long as all the links match up and the `x-kubernetes-group-version-kind`
	// metadata is correct
	var kruiseMappings = map[string]string{
		"github.com/openkruise/kruise-api/apps":   "apps.kruise.io",
		"github.com/openkruise/kruise-api/policy": "policy.kruise.io",
	}

	defMap := GetOpenAPIDefinitions(func(path string) spec.Ref {
		for k, v := range kruiseMappings {
			path = strings.ReplaceAll(path, k, v)
		}
		return spec.MustCreateRef(fmt.Sprintf("#/definitions/%s", kubeopenapiutil.ToRESTFriendlyName(path)))
	})

	var crdDefs = make(map[string]openAPISchema)
	for pathKey, definition := range defMap {
		for k, v := range kruiseMappings {
			pathKey = strings.ReplaceAll(pathKey, k, v)
		}
		crdDefs[kubeopenapiutil.ToRESTFriendlyName(pathKey)] = openAPISchema{
			Schema:                      definition.Schema,
			XKubernetesGroupVersionKind: make([]xKubernetesGroupVersionKind, 0),
		}
	}

	k8sDefs, err := loadK8SDefinitions()
	checkErr(err)

	for k, v := range crdDefs {
		//We pull out openkruise crd information based on the dot pattern of the key in the dictionary we are also setting it for all
		//argo types instead of just the ones needed this could be incorrect as far as spec goes, but it works.
		if strings.HasPrefix(k, "io.kruise") {
			kruiseGVK := strings.Split(k, ".")
			v.XKubernetesGroupVersionKind = []xKubernetesGroupVersionKind{
				{
					Group:   fmt.Sprintf("%s.kruise.io", kruiseGVK[2]),
					Kind:    kruiseGVK[4],
					Version: kruiseGVK[3],
				},
			}
			crdDefs[k] = v
			continue
		}

		// Pull the group version kind information from the k8s definitions that downloaded via loadK8SDefinitions
		if entry, ok := k8sDefs.Definitions[k]; ok {
			if len(entry.XKubernetesGroupVersionKind) > 0 {
				v.XKubernetesGroupVersionKind = entry.XKubernetesGroupVersionKind
				crdDefs[k] = v
			}
		}
	}

	data, err := json.MarshalIndent(map[string]interface{}{
		"definitions": crdDefs,
	}, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, data, 0644)
}

// loadK8SDefinitions loads K8S types API schema definitions starting with the version specified in go.mod then the fucnction
// parameter versions
func loadK8SDefinitions() (*k8sGvkMapping, error) {
	// detects minor version of k8s client
	k8sVersionCmd := exec.Command("sh", "-c", "cat go.mod | grep \"k8s.io/client-go\" |  head -n 1 | cut -d' ' -f2")
	versionData, err := k8sVersionCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to determine k8s client version: %v", err)
	}
	v, err := semver.Parse(strings.TrimSpace(strings.Replace(string(versionData), "v", "", 1)))
	if err != nil {
		return nil, err
	}

	schemaGoMod := k8sGvkMapping{}
	err = downloadK8SDefinitions(uint(v.Minor), &schemaGoMod)
	if err != nil {
		return nil, err
	}

	for _, v := range K8SVersions {
		// Download fixe old version to keep old schema's compatibility
		schemaFixedVer := k8sGvkMapping{}
		err = downloadK8SDefinitions(uint(v), &schemaFixedVer)
		if err != nil {
			return nil, err
		}

		// Merge old and new schema
		for k, v := range schemaFixedVer.Definitions {
			schemaGoMod.Definitions[k] = v
		}
	}

	return &schemaGoMod, nil
}

func downloadK8SDefinitions(version uint, k8sGvkMapping *k8sGvkMapping) error {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/release-1.%d/api/openapi-spec/swagger.json", version))
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, k8sGvkMapping)
	if err != nil {
		return err
	}

	return nil
}

// Generate CRD spec for Resources in OpenKruise
func main() {
	err := generateOpenApiSchema("schema/openkruise_kruise_kustomize_schema.json", kruise.GetOpenAPIDefinitions)
	checkErr(err)
	err = generateOpenApiSchema("schema/openkruise_all_k8s_kustomize_schema.json", apis.GetOpenAPIDefinitions)
	checkErr(err)
}
