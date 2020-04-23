package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	gw "nice/hello"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		return err
	}
	r := http.NewServeMux()
	r.Handle("/", mux)
	log.Println("listening to port *:8080")
	return http.ListenAndServe(":8080", r)
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
