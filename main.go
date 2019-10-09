package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ehlxr/goproxy/metadata"
	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
	"github.com/mitchellh/go-homedir"
)

var (
	port    = flag.Int("port", 8080, "the port of goproxy server")
	host    = flag.String("host", "0.0.0.0", "the host of goproxy server")
	version = flag.Bool("version", false, "Show version info")
)

type MyGoPoroxy struct {
	*goproxy.Goproxy
}

func main() {
	printVersion()

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		// _, _ = fmt.Fprintf(os.Stdout, "Usage of %s:\n", "goproxy")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *version {
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	if strings.Contains(addr, "0.0.0.0") {
		addr = strings.Replace(addr, "0.0.0.0", "", 1)
		*host = strings.Replace(*host, "0.0.0.0", "127.0.0.1", 1)
	}
	fmt.Printf("goproxy server start on: %s\n", fmt.Sprintf("http://%s:%d", *host, *port))

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gp := goproxy.New()
	gp.Cacher = &cacher.Disk{Root: fmt.Sprintf("%s/.goproxy", home)}

	if err := http.ListenAndServe(addr, &MyGoPoroxy{gp}); err != nil {
		log.Fatalf("goproxy server error: %v", err)
	}
}

func (g *MyGoPoroxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	trimmedPath := path.Clean(r.URL.Path)
	trimmedPath = strings.TrimPrefix(trimmedPath, g.PathPrefix)
	trimmedPath = strings.TrimLeft(trimmedPath, "/")

	if trimmedPath == "" {
		_, _ = rw.Write([]byte("go proxy server is running!"))

		return
	}

	g.Goproxy.ServeHTTP(rw, r)
}

// printVersion Print out version information
func printVersion() {
	banner, _ := base64.StdEncoding.DecodeString(metadata.BannerBase64)
	fmt.Printf(metadata.VersionTpl, banner, metadata.Version, metadata.BuildTime, metadata.GitCommit, metadata.GoVersion)
}
