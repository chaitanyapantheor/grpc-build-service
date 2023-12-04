package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chaitanyapantheor/grpc-build-service/build"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := build.NewBuildServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBuilds(ctx, &build.BuildRequest{Filter: "all"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("GetBuilds: %s", r.GetBuild().GetLabel())
}
