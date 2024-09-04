#!/usr/bin/env bash

go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -e
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/github.com/openkruise/kruise-api
cp -r ./{apps,policy,hack,vendor,go.mod,.git} "${TMP_DIR}"/src/github.com/openkruise/kruise-api/
chmod +x "${TMP_DIR}"/src/github.com/openkruise/kruise-api/vendor/k8s.io/code-generator/generate-internal-groups.sh
echo "tmp_dir: ${TMP_DIR}"

(cd "${TMP_DIR}"/src/github.com/openkruise/kruise-api; \
    GOPATH=${TMP_DIR} GO111MODULE=on /bin/bash vendor/k8s.io/code-generator/generate-groups.sh client,deepcopy,informer,lister \
    github.com/openkruise/kruise-api/client github.com/openkruise/kruise-api "apps:v1alpha1 apps:v1beta1 policy:v1alpha1" -h ./hack/boilerplate.go.txt)

mkdir -p ./client
rm -rf ./client/{clientset,informers,listers}
mv "${TMP_DIR}"/src/github.com/openkruise/kruise-api/client/* ./client/
