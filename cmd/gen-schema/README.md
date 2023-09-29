## gen-schema
### Description
This dictionary contains a tool that generates a kustomize schema file 
for native kubernetes resources plus projects under the openkrusie name including kruise and rollouts.

The purpose of the kustomize schema file is to be able to use Policy Merge Patch (SMP) 
on any array type field within an environment variable or resource definition 
when using kustomize to deploy CRDs in openkruise.
### How to run
To generate the schema, run the following command in the project root directory:
```
make gen-openapi-schema
```
the schema will be generated in the `schema` folder.

### How to use 
In your kustomization.yaml file add a config block like:

```
openapi:
  path: https://raw.githubusercontent.com/openkruise/kruise-api/main/schema/openkruise_all_k8s_kustomize_schema.json
```

you can find examples in `test` folder.

### How to upgrade the schema
When there are some new CRDs or updated CRD fields in openkruise, 
some new fields of array types may need to support SMP.
You can simply upgrade schema in the following steps:

#### 1. Check field and add annotations
Firstly, determine whether the newly added field is of array type.
Then, combined with the usage scenario, determine if there is a requirement 
for merging the field based on a unique **key**.

Then, add annotations to the specified field in the CRD definition file as follows:
```
// +patchMergeKey=name
// +patchStrategy=merge
InitContainers []SidecarContainer `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name"`
```

#### 2. Run command to generate schema
```
make gen-openapi-schema
```

### Related discussions and issues
1. [Why include native resources definition info the schema file?](https://github.com/argoproj/argo-rollouts/issues/1730)
2. [Kustomize build not honoring the retainKeys patchStrategy](https://github.com/kubernetes-sigs/kustomize/issues/3981)
3. [Reason for Error `API rule violation: list_type_missing`](https://github.com/kubernetes/kube-openapi/issues/175#issuecomment-546504117)
