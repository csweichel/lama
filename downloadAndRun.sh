#!/bin/sh

#
# This script downloads and starts servr: a simple HTTP server that serves files, directories and verbosely logs requests.
#
# run with "curl -L servr.world | sh"
# 

export platform="$(uname -s)_$(uname -m)"
export servr=$HOME/.servr

export download=true
if [ ! -f $servr ]; then
    download=true
elif [ `which sha256sum` ]; then
    download=false

    export latestChecksum=$(curl --silent -s https://api.github.com/repos/32leaves/servr/releases/latest \
                    | grep browser_download_url \
                    | grep checksums.txt \
                    | cut -d : -f 2,3 \
                    | xargs curl --silent -L \
                    | grep $platform \
                    | cut -d ' ' -f 1)
    currentChecksum=$(sha256sum $servr | cut -d ' ' -f 1)

    if [ $latestChecksum != $currentChecksum ]; then
        download=true
    fi
else
    echo "No sha256sum found - skipping checksum checks"
    download=true
fi

if [ "$download" = "true" ]; then
    echo "Downloading latest servr version"
    curl --silent -s https://api.github.com/repos/32leaves/servr/releases/latest \
    | grep browser_download_url \
    | grep $platform \
    | cut -d : -f 2,3 \
    | xargs curl --silent -L --output $servr

    chmod +x $servr
fi

# at least we can start servr itself
exec $servr $@
