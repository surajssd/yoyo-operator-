#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/surajssd/operators/using-operator-sdk/yoyo-operator/pkg/generated \
github.com/surajssd/operators/using-operator-sdk/yoyo-operator/pkg/apis \
yoyo:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
