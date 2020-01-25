package  main

import (
	"context"
	pb "github.com/ArtGooner/test-microservice/config/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.Account) (*pb.User, error) {
	log.Printf("Email: %v", in.GetEmail())
	log.Printf("PasswordHash: %v", in.GetPasswordHash())
	return &pb.User{Id: 1, Name: "Artsiom", Surname: "Ivaniutsenka", Age: 27}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
