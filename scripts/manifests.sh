#!/bin/bash
set -e
kustomize build kustomize/overlays/integration/ -o k8s/integration/manifests.yaml
kustomize build kustomize/overlays/prod/ -o k8s/prod/manifests.yaml
