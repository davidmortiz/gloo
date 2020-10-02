#!/bin/bash

cd ../
git clone https://github.com/solo-io/solo-apis.git
cd solo-apis
git config --global user.name "Gloo Github Action"
./hack/sync-gloo-apis.sh; make update-deps generate -B
git checkout -b update-solo-apis
git add .
git commit -m "update to latest gloo version"
git push origin master
