#!/bin/bash

if [ "$GITPOD_WORKSPACE_ID" != "" ]; then
    export PATH=/workspace:$PATH
fi

find dist -type f -name lama | grep linux | xargs upx --brute $i