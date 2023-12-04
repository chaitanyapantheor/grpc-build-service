package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/chaitanyapantheor/grpc-build-service/backend"
	"github.com/chaitanyapantheor/grpc-build-service/build"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func run() {
	//run grpc server
	s := grpc.NewServer()
	build.RegisterBuildServiceServer(s, &backend.Server{})
	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.Serve(lis)
}

func main() {
	flag.Parse()
	go run()

	conn, err := grpc.DialContext(
		context.Background(),
		*grpcServerEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return
	}

	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx := context.Background()
	// err = build.RegisterBuildServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	err = build.RegisterBuildServiceHandler(ctx, mux, conn)
	if err != nil {
		return
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8081", mux)

}
