#!/usr/bin/env bash

# Copyright 2018 The Knative Authors
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

source $(dirname $0)/../vendor/github.com/knative/test-infra/scripts/library.sh

# Remove symlinks in /vendor that are broken or lead outside the repo.
function remove_broken_symlinks() {
  for link in $(find ./vendor -type l); do
    # Remove broken symlinks
    if [[ ! -e ${link} ]]; then
      unlink ${link}
      continue
    fi
    # Get canonical path to target, remove if outside the repo
    local target="$(ls -l ${link})"
    target="${target##* -> }"
    [[ ${target} == /* ]] || target="./${target}"
    target="$(cd `dirname ${link}` && cd ${target%/*} && echo $PWD/${target##*/})"
    if [[ ${target} != *github.com/projectriff/* ]]; then
      unlink ${link}
      continue
    fi
  done
}

cd ${REPO_ROOT_DIR}

# Ensure we have everything we need under vendor/
dep ensure

rm -rf $(find vendor/ -name 'OWNERS')
rm -rf $(find vendor/ -name '*_test.go')

update_licenses third_party/VENDOR-LICENSE "./cmd/*"

# Patch the Kubernetes dynamic client to fix listing. This patch is from
# https://github.com/kubernetes/kubernetes/pull/68552/files, which is a
# cherrypick of #66078.  Remove this once that reaches a client version
# we have pulled in.
git apply ${REPO_ROOT_DIR}/hack/66078.patch

# Remove all invalid symlinks under ./vendor
remove_broken_symlinks