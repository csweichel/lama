package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"net/http"
	"net/http/httputil"
)

var build = "UNKNOWN"

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	port := pflag.StringP("port", "p", "8080", "port to serve on")
	dir := pflag.StringP("directory", "d", ".", "the directory to serve")
	local := pflag.BoolP("localhost", "l", false, "serve on localhost only")
	version := pflag.BoolP("version", "v", false, "prints the version")
	dontDump := pflag.BoolP("dont-dump", "N", false, "be less verbose and don't dump requests")
	dontServe := pflag.BoolP("dont-serve", "D", false, "don't serve any directy (ignores --directory)")
	pflag.Parse()

	if *version {
		fmt.Println(build)
		return
	}

	handler := &debugHandler{DumpRequest: !*dontDump}
	if !*dontServe {
		log.WithField("directory", *dir).Info("serving files")
		handler.Delegate = http.FileServer(http.Dir(*dir))
	}
	http.Handle("/", handler)

	addr := ":" + *port
	if *local {
		addr = "127.0.0.1" + addr
	}

	log.WithField("addr", addr).WithField("version", build).Info("server running")
	err := http.ListenAndServe(addr, nil)
	log.WithError(err).Fatal("cannot serve")
}

type debugHandler struct {
	Delegate    http.Handler
	DumpRequest bool
}

func (h *debugHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	dump := fmt.Sprintf("%s %s", req.Method, req.URL)
	if h.DumpRequest {
		out, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.WithError(err).Error("cannot dump request")
		} else {
			dump = string(out)
		}
	}
	log.Info(dump)

	if h.Delegate != nil {
		h.Delegate.ServeHTTP(resp, req)
	}
}
