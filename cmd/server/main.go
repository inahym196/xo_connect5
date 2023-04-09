package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "github.com/inahym196/xo_connect5/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGameServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Body: fmt.Sprintf("Hello, %s!", req.GetBody()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hellopb.RegisterGameServiceServer(s, NewMyServer())
	reflection.Register(s)
	go func() {
		log.Printf("start gRPC server port:%v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gPRC server...")
	s.GracefulStop()
}
