package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "nice/hello"
)

type helloServer struct {
}

func (s *helloServer) Echo(ctx context.Context, msg *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	log.Println(msg, ctx)
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})
	log.Println("listening to port *:9090")
	grpcServer.Serve(lis)
}
