package main

import (
	"context"
	"github.com/ArtGooner/test-microservice/auth/user"
	pb "github.com/ArtGooner/test-microservice/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.Account) (*pb.User, error) {
	rps, err := user.NewRepository()

	if err != nil {
		log.Fatal(err)
	}

	res, err := rps.Get(in)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return &pb.User{Id: res.Id, Name: res.Name, Surname: res.Surname, Age: res.Age, PasswordHash: res.PasswordHash}, nil
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
