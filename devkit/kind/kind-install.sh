#!/bin/bash

KIND_VERSION=0.14.0

curl -Lo ./kind "https://kind.sigs.k8s.io/dl/v${KIND_VERSION}/kind-$(uname)-amd64"
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
