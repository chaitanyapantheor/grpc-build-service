package backend

import "github.com/chaitanyapantheor/grpc-build-service/build"

// server is used to implement helloworld.GreeterServer.
type Server struct {
	build.UnimplementedBuildServiceServer
}
