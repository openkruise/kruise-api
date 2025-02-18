#!/usr/bin/env bash

go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
	    exit $retVal
fi

set -e
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/github.com/openkruise/kruise-api
cp -r ./{apps,policy,hack,utils,vendor,go.mod,.git} "${TMP_DIR}"/src/github.com/openkruise/kruise-api/
echo "tmp_dir: ${TMP_DIR}"
SCRIPT_ROOT="${TMP_DIR}"/src/github.com/openkruise/kruise-api
CODEGEN_PKG=${CODEGEN_PKG:-"${SCRIPT_ROOT}/vendor/k8s.io/code-generator"}
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "gen_openapi"
report_filename="./violation_exceptions.list"

(cd "${TMP_DIR}"/src/github.com/openkruise/kruise-api; \
	GOPATH=${TMP_DIR} kube::codegen::gen_openapi \
	    --output-dir "${SCRIPT_ROOT}/pkg/kruise" \
	    --output-pkg "github.com/openkruise/kruise-api/pkg/kruise" \
	    --extra-pkgs "github.com/openkruise/kruise-api/policy/v1alpha1" \
            --report-filename "${report_filename}" \
            --update-report \
	    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
	    "${SCRIPT_ROOT}/apps" 
            )

(cd "${TMP_DIR}"/src/github.com/openkruise/kruise-api; \
        GOPATH=${TMP_DIR} kube::codegen::gen_openapi \
            --output-dir "${SCRIPT_ROOT}/pkg/apis" \
            --output-pkg "github.com/openkruise/kruise-api/pkg/apis" \
	    --extra-pkgs "github.com/openkruise/kruise-api/policy/v1alpha1" \
	    --extra-pkgs "k8s.io/api/autoscaling/v1" \
	    --extra-pkgs "k8s.io/api/batch/v1" \
	    --extra-pkgs "k8s.io/api/certificates/v1" \
	    --extra-pkgs "k8s.io/api/core/v1" \
	    --extra-pkgs "k8s.io/api/networking/v1" \
	    --extra-pkgs "k8s.io/api/policy/v1" \
	    --extra-pkgs "k8s.io/api/rbac/v1" \
	    --extra-pkgs "k8s.io/api/storage/v1" \
            --report-filename "${report_filename}" \
            --update-report \
            --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
            "${SCRIPT_ROOT}/apps"
            )

cp -f "${TMP_DIR}"/src/github.com/openkruise/kruise-api/pkg/kruise/zz_generated.openapi.go ./pkg/kruise/openapi_generated.go
cp -f "${TMP_DIR}"/src/github.com/openkruise/kruise-api/pkg/apis/zz_generated.openapi.go ./pkg/apis/openapi_generated.go
rm -rf vendor 
rm -rf ${TMP_DIR}
