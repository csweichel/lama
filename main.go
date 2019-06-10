// Copyright (c) 2019 Christian Weichel

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

// version set during build
var build = "UNKNOWN"

var (
	fgURL    = color.New(color.FgHiWhite, color.Bold)
	fgHeader = color.New(color.FgHiYellow)
	fgMethod = color.New(color.FgCyan)
	fgError  = color.New(color.FgHiRed, color.Bold)
	fgLama   = color.New(color.FgHiCyan, color.Bold)
)

func main() {
	port := pflag.StringP("port", "p", "8080", "port to serve on")
	dir := pflag.StringP("directory", "d", ".", "the directory to serve")
	local := pflag.BoolP("localhost", "l", false, "serve on localhost only")
	version := pflag.BoolP("version", "v", false, "prints the version")
	verbose := pflag.BoolP("dump", "V", false, "be verbose and dump requests")
	dontServe := pflag.BoolP("dont-serve", "D", false, "don't serve any directy (ignores --directory)")
	pflag.Parse()

	if *version {
		fmt.Println(build)
		return
	}

	handler := &debugHandler{
		DumpRequest: *verbose,
		Writer:      tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', 0),
		Logger:      make(chan *http.Request, 10),
	}
	go handler.logRequests()

	fileStatement := ""
	if !*dontServe {
		handler.Delegate = http.FileServer(http.Dir(*dir))
		fileStatement = fmt.Sprintf("files from %s ", *dir)
	}
	http.Handle("/", handler)

	addr := ":" + *port
	if *local {
		addr = "127.0.0.1" + addr
	}

	fmt.Printf("This is %s serving %son %s\r\n", fgLama.Sprint("lama.sh"), fileStatement, addr)
	if !*verbose {
		go cli(handler)
	}

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("%s %s - %s\n", fgError.Sprint("ERROR"), time.Now().Format(time.RFC3339), err.Error())
		os.Exit(1)
	}
}

type debugHandler struct {
	Delegate    http.Handler
	DumpRequest bool
	Writer      *tabwriter.Writer
	Logger      chan *http.Request
}

func (h *debugHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.Logger <- req

	if h.Delegate != nil {
		h.Delegate.ServeHTTP(resp, req)
	}
}

func (h *debugHandler) logRequests() {
	for {
		req := <-h.Logger
		fmt.Printf("%s %s - %s %s\r\n",
			fgMethod.Sprintf("%-7v", req.Method),
			time.Now().Format(time.RFC3339),
			req.Proto,
			fgURL.Sprint(req.URL),
		)
		if !h.DumpRequest {
			continue
		}

		headerPadding := fmt.Sprintf("%-8v", " ")

		// from https://golang.org/src/net/http/httputil/dump.go?s=5638:5700#L219
		absRequestURI := strings.HasPrefix(req.RequestURI, "http://") || strings.HasPrefix(req.RequestURI, "https://")
		if !absRequestURI {
			host := req.Host
			if host == "" && req.URL != nil {
				host = req.URL.Host
			}
			if host != "" {
				fmt.Fprintf(h.Writer, "%s%s\t%s\r\n", headerPadding, fgHeader.Sprint("Host:"), host)
			}
		}

		if len(req.TransferEncoding) > 0 {
			fmt.Fprintf(h.Writer, "%s%s\t%s\r\n", headerPadding, fgHeader.Sprint("Transfer-Encoding:"), strings.Join(req.TransferEncoding, ","))
		}
		if req.Close {
			fmt.Fprintf(h.Writer, "%s%s\tclose\r\n", headerPadding, fgHeader.Sprint("Connection:"))
		}

		for k, v := range req.Header {
			fmt.Fprintf(h.Writer, "%s%s\t%s\r\n", headerPadding, fgHeader.Sprintf("%s:", k), strings.Join(v, ", "))
		}

		h.Writer.Flush()
		fmt.Println()
	}
}

func cli(handler *debugHandler) {
	in := bufio.NewReader(os.Stdin)

	for {
		if runtime.GOOS == "darwin" {
			// On OSX we'll eat the remainder of the download script which messes with this ReadLine mechanism.
			// For now, we'll just always enable verbose logging on OSX.
		} else {
			fmt.Print("press <enter> to enable verbose log output\r\n\r\n")
			_, _, err := in.ReadLine()
			if err != nil {
				continue
			}
		}

		handler.DumpRequest = true
		fmt.Printf("%s %s - verbose logging enabled\r\n", fgLama.Sprint("lama.sh"), time.Now().Format(time.RFC3339))
		break
	}
}
