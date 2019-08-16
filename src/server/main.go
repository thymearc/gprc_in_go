//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main


import (
	"context"
	"log"
	"fmt"
	"strconv"
	"net"
	"google.golang.org/grpc"
	pb "./CalcService"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) CalcAdd(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	a, err := strconv.Atoi(in.A)
	if err != nil {
		log.Fatalf("failed to convert: %v", err)
	}
	b, err := strconv.Atoi(in.B)
	if err != nil {
		log.Fatalf("failed to convert: %v", err)
	}
	return &pb.Reply{fmt.Sprintf("result: %d", ( a + b ))}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
