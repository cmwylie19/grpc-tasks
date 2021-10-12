package main

import (
	"log"
	"net"

	"github.com/cmwylie19/knative-poc/api"
	"github.com/cmwylie19/knative-poc/controllers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterTodoServer(s, &controllers.Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
