package backend

import (
	"context"
	"log"

	"github.com/chaitanyapantheor/grpc-build-service/build"
	"github.com/google/uuid"
)

// GetBuilds implements build.BuildServiceServer
func (s *Server) GetBuilds(ctx context.Context, in *build.BuildRequest) (*build.BuildResponse, error) {
	log.Printf("Received: %v", in.GetFilter())
	return &build.BuildResponse{Build: &build.Build{Id: uuid.New().String(), Label: "Build 1", Status: "Pending"}}, nil
}
