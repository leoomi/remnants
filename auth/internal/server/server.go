package server

import (
	"context"
	"log"
	"net"

	"github.com/leoomi/remnants/auth/internal/config"
	"github.com/leoomi/remnants/auth/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthenticationServer
}

func (s *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.VerifyResponse{}, nil
}

func RunServer() {
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %s", err)
	}

	lis, err := net.Listen("tcp", c.AuthServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
