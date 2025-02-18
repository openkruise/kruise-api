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

chmod +x "${TMP_DIR}"/src/github.com/openkruise/kruise-api/vendor/k8s.io/code-generator/generate-internal-groups.sh
echo "tmp_dir: ${TMP_DIR}"

SCRIPT_ROOT="${TMP_DIR}"/src/github.com/openkruise/kruise-api
CODEGEN_PKG=${CODEGEN_PKG:-"${SCRIPT_ROOT}/vendor/k8s.io/code-generator"}

echo "source ${CODEGEN_PKG}/kube_codegen.sh"
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "gen_helpers"
GOPATH=${TMP_DIR} GO111MODULE=off kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}"

echo "gen_client"
GOPATH=${TMP_DIR} GO111MODULE=off kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/client" \
    --output-pkg "github.com/openkruise/kruise-api/client" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}"

mkdir -p ./client
rm -rf ./client/{clientset,informers,listers}
mv "${TMP_DIR}"/src/github.com/openkruise/kruise-api/client/* ./client/
