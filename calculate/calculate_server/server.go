package main

import (
	"context"
	"fmt"
	calculatepb "go-grpc/calculate/calculatepb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

// Sum ...
func (*server) Sum(ctx context.Context, req *calculatepb.SumRequest) (*calculatepb.SumResponse, error) {
	fmt.Println("Recivied Sum RPC ", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatepb.SumResponse{
		SumResult: sum,
	}
	return res, nil

}

func main() {
	fmt.Println("Calculate server")
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server", err)
	}
}
