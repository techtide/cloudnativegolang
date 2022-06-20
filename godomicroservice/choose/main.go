//go:generate protoc -I . --go_out=plugins=grpc:. choose.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"godomicroservice/pb"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Couldn't listen because of error %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChooseServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Couldn't serve because of error %v", err)
	}
}

func factorial(x float64) float64 {
	if x >= 1 {
		return x * factorial(x-1)
	}
	return 1
}

func choose(n, k float64) float64 {
	return (factorial(n)) / (factorial(k) * factorial(n-k))
}

func (s *server) Compute(cxt context.Context, r *pb.ChooseRequest) (*pb.ChooseResponse, error) {
	result := &pb.ChooseResponse{}
	result.Result = choose(r.N, r.K)

	logMessage := fmt.Sprintf("n: %f   k: %f     nCk: %f", r.N, r.K, result.Result)
	log.Println(logMessage)

	return result, nil
}
