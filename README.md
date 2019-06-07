# Getting started
```
curl -L servr.world | sh
```

## Configuration
`server` sports a few command-line flags which configure its behaviour. To pass those flags use
```
$ curl -L servr.world | sh -s -- <flags>
```
For example to print the list of available flags
```
$ curl -L servr.world | sh -s -- --help
Usage of servr:
  -d, --directory string   the directory to serve (default ".")
  -N, --dont-dump          be less verbose and don't dump requests
  -D, --dont-serve         don't serve any directy (ignores --directory)
  -l, --localhost          serve on localhost only
  -p, --port string        port to serve on (default "8080")
  -v, --version            prints the version
```

# Introduction
Servr is a simple HTTP server that serves files, directories and verbosely logs requests.
This project has two main qualities:
1. **Easy to use:** servr allows one to start a web server on any Linux/MacOSX using 25 characters, with no requirements (not even libc).
2. **Trustworthy:** servr's code and download script are as simple as they can be, thus can be inspected and verified. We have a very shallow
   dependency tree (just two dependencies other than the Go standard library).

# How to contribute
All contributions/PR/issue/beer are welcome ❤️.

It's easiest to develop _servr_ using Gitpod, a free one-click online IDE (who I happen to be working on):

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io#github.com/32leaves/servr)
