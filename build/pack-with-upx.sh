#!/bin/bash

set -e

if [ "$GITPOD_WORKSPACE_ID" != "" ]; then
    export PATH=/workspace:$PATH
fi

# Check if dist directory exists
if [ ! -d "dist" ]; then
    echo "ERROR: dist directory not found. Run build first."
    exit 1
fi

# Check if upx is available
if ! command -v upx >/dev/null 2>&1; then
    echo "ERROR: upx command not found. Install upx first."
    exit 1
fi

# Find matching files
files=$(find dist -type f -name lama | grep linux | grep -v arm64)

if [ -z "$files" ]; then
    echo "ERROR: No matching linux lama binaries found in dist directory."
    exit 1
fi

echo "Compressing binaries with upx:"
echo "$files" | xargs upx --brute
