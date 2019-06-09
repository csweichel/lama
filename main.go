package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"
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
	dontDump := pflag.BoolP("dont-dump", "N", false, "be less verbose and don't dump requests")
	dontServe := pflag.BoolP("dont-serve", "D", false, "don't serve any directy (ignores --directory)")
	pflag.Parse()

	if *version {
		fmt.Println(build)
		return
	}

	handler := &debugHandler{
		DumpRequest: !*dontDump,
		Writer:      tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', 0),
	}

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

	fmt.Printf("This is %s serving %son %s\r\n\r\n", fgLama.Sprint("lama.sh"), fileStatement, addr)

	err := http.ListenAndServe(addr, nil)
	logError("cannot serve", err)
}

type debugHandler struct {
	Delegate    http.Handler
	DumpRequest bool
	Writer      *tabwriter.Writer
}

func (h *debugHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.logRequest(req, h.DumpRequest)

	if h.Delegate != nil {
		h.Delegate.ServeHTTP(resp, req)
	}
}

func logError(msg string, err error) {
	fgError.Print("ERROR")
	fmt.Printf(" %s - %s\n", time.Now().Format(time.RFC3339), err.Error())
}

func (h *debugHandler) logRequest(req *http.Request, verbose bool) {
	fmt.Printf("%s %s - %s %s\r\n",
		fgMethod.Sprintf("%-7v", req.Method),
		time.Now().Format(time.RFC3339),
		req.Proto,
		fgURL.Sprint(req.URL),
	)
	if !verbose {
		return
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
