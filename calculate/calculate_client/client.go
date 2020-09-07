package main

import (
	"context"
	"fmt"
	"go-grpc/calculate/calculatepb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("calculate client")
	conn, err := grpc.Dial("127.0.1.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer conn.Close()
	c := calculatepb.NewCalculateServiceClient(conn)
	doUnary(c)
}

func doUnary(c calculatepb.CalculateServiceClient) {
	fmt.Println("Starting to sum ...")
	req := &calculatepb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("err", err)
	}
	log.Println("result", res.SumResult)

}
