#!/bin/sh

# <div style="display:none">
# Copyright (c) 2019 Christian Weichel
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
# </div>

# <link href="style.css" rel="stylesheet"><div id="m"><div id="c"><img id="l" src="lama.svg"><div><h2>curl lama.sh | sh</h2><h3>starts a web server</h3><p><a href="https://github.com/csweichel/lama">lama.sh</a> is a simple HTTP server that serves files, directories and verbosely logs requests. Running the line above will start an HTTP server. Wanna know how? Have a look at <a href="https://github.com/csweichel/lama/blob/master/docs/index.html">this page's source</a>.</p><div id="b"><a class="btn" href="https://github.com/csweichel/lama"><div><img src="github-logo.svg" height="20px"> View on GitHub</div></a><a href="https://gitpod.io/#github.com/csweichel/lama"><img src="https://gitpod.io/button/open-in-gitpod.svg"></a></div></div></div></div><div id="h">

#
# This script downloads and starts a simple HTTP server that serves files, directories and verbosely logs requests.
#
# run with "curl lama.sh | sh"
#

export platform="$(uname -s)_$(uname -m)"
export lama=$HOME/.lama

latestRelease=$(curl -L --silent -s https://api.github.com/repos/csweichel/lama/releases/latest \
    | grep browser_download_url \
    | grep $platform \
    | cut -d : -f 2,3)

export download=true
if [ -f $lama ]; then
    installedVersion=$($lama -v)

    if echo $latestRelease | grep $installedVersion >/dev/null; then
        download=false
    else
        download=true
    fi
fi

if [ "$download" = "true" ]; then
    echo "Downloading latest lama version from$latestRelease"
    echo $latestRelease | xargs curl -L --output $lama

    chmod +x $lama

    if [ `which sha256sum` ]; then
        export latestChecksum=$(curl --silent -s https://api.github.com/repos/csweichel/lama/releases/latest \
                        | grep browser_download_url \
                        | grep checksums.txt \
                        | cut -d : -f 2,3 \
                        | xargs curl --silent -s -L \
                        | grep $platform \
                        | cut -d ' ' -f 1)
        currentChecksum=$(sha256sum $lama | cut -d ' ' -f 1)

        if [ $latestChecksum != $currentChecksum ]; then
            echo "checksum mismatch - will not start downloaded binary"
            rm $lama
            exit -1
        fi
    fi
fi

# at least we can start lama itself
exec $lama $@

#
# Why is this called index.html?
#   I want to re-use GitHub infrastructure, specifically GitHub pages where the index file must
#   be named index.html.
#
# What's with the pre tags?
#   Because this file is called index.html, GitHub will serve it with an HTML content-type, which
#   causes the browser to render this script funny. But I want to be easy to read so that people
#   can understand what's going on here.
#
# The last line has to be </div> as on OSX we're seeing the remainer of this script and need to filter that.
# </div>
