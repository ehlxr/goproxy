package main

import (
	"flag"
	"fmt"
	"github.com/ehlxr/goproxy/util"
	"github.com/goproxy/goproxy"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port    = flag.Int("port", 8080, "the port of goproxy server")
	host    = flag.String("host", "0.0.0.0", "the host of goproxy server")
	version = flag.Bool("version", false, "Show version info")
)

func main() {

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		// _, _ = fmt.Fprintf(os.Stdout, "Usage of %s:\n", "goproxy")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *version {
		util.PrintVersion()
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	if strings.Contains(addr, "0.0.0.0") {
		addr = strings.Replace(addr, "0.0.0.0", "", 1)
		*host = strings.Replace(*host, "0.0.0.0", "127.0.0.1", 1)
	}
	fmt.Printf("goproxy server start on: %s\n", fmt.Sprintf("http://%s:%d", *host, *port))

	if err := http.ListenAndServe(addr, goproxy.New()); err != nil {
		log.Fatalf("goproxy server error: %v", err)
	}
}
