#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=/home/lidongqi/work/k8s/k8s.io/code-generator

source "${CODEGEN_PKG}/kube_codegen.sh"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
export KUBE_VERBOSE=2

#-i 是一个package路径，不要最后的 '/'
register-gen -v ${KUBE_VERBOSE} -i github.com/lidongqi/demo-controller/apis/ldq.test.com/v1 \
--go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
-o "$(dirname "${BASH_SOURCE[0]}")/../../../.."

kube::codegen::gen_helpers \
    --input-pkg-root github.com/lidongqi/demo-controller/apis \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root github.com/lidongqi/demo-controller/apis \
    --output-pkg-root github.com/lidongqi/demo-controller/generated \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
