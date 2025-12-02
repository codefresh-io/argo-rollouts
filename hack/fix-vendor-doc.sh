#!/bin/bash

# This script fixes a compatibility issue between controller-gen v0.14.0 and
# the vendored k8s.io/api core/v1 package. The +groupName= annotation needs
# a value (empty string for core API group) to be properly parsed.

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "${SCRIPT_DIR}/.." && pwd )"

echo "Fixing vendored doc.go files for controller-gen compatibility..."

# Fix k8s.io/api/core/v1/doc.go
sed -i 's|^// +groupName=$|// +groupName=""|' "${PROJECT_ROOT}/vendor/k8s.io/api/core/v1/doc.go"

# Fix k8s.io/kubernetes/pkg/apis/core/doc.go
sed -i 's|^// +groupName=$|// +groupName=""|' "${PROJECT_ROOT}/vendor/k8s.io/kubernetes/pkg/apis/core/doc.go"

echo "Vendor doc.go files fixed successfully"

