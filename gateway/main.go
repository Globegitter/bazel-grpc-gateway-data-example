package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/globegitter/bazel-grc-gateway-data-example/proto/service"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:8080", "endpoint of YourService")
	swaggerDir   = flag.String("swagger_dir", "gateway", "path to the directory which contains swagger definitions")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwMux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterBillingHandlerFromEndpoint(ctx, gwMux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	mux.Handle("/", gwMux)

	return http.ListenAndServe(":9090", mux)
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received call to swagger: %s", r.URL.Path)
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	p = path.Join(dir, "gateway.runfiles", "__main__", *swaggerDir, p)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Serving %s from %s", r.URL.Path, p)
	http.ServeFile(w, r, p)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
