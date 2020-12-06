package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
	"io"
	"log"
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().Unix())
}

func main()  {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCounterClient(conn)

	r, err := c.Sum(context.Background())

	for {
		req := &pb.NumberReq{X: rand.Int63n(100), Y: rand.Int63n(100)}
		err := r.Send(req)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		fmt.Printf("Client Sended: [X = %d, Y = %d]\n", req.X, req.Y)
		resp, err := r.Recv()
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		fmt.Printf("Client Received: [Result = %d]\n", resp.Result)
		time.Sleep(time.Second)
	}
}
