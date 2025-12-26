#!/usr/bin/env bash

[ -d kruise ] || git clone git@github.com:openkruise/kruise.git
mkdir -p output
cp -r kruise/apis/* output/
# do not expose default package
rm -rf output/apps/defaults
find output -type f -exec sed -i 's%github.com/openkruise/kruise/apis/apps%github.com/openkruise/kruise-api/apps%g' {} +
find output -type f -exec sed -i 's%github.com/openkruise/kruise/apis/policy%github.com/openkruise/kruise-api/policy%g' {} +
cp -r output/* ./
