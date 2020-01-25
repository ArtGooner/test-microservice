package main

import (
	"fmt"
	pb "github.com/ArtGooner/test-microservice/config"
	"github.com/ArtGooner/test-microservice/handler"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
	port    = "50051"
)

func main() {
	/* grpc Setup */
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	/* Echo Setup*/
	e := echo.New()
	e.HideBanner = true

	eh := handler.New(c)
	e.GET("/Login", eh.Login)
	//e.GET("/Register", eh.Create)
	//e.GET("/Get", eh.Get)

	err = e.Start(fmt.Sprintf("HTTP Server started on port:%d", port))

	if err != nil {
		log.Fatal(err)
	}
}
