package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/proto"
	"io"
	"log"
	"net"
)

const Port = 8080

type server struct {
	pb.CounterServer
}

func (s *server) Sum(stream pb.Counter_SumServer) error {
	for  {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Server Received: [X = %d, Y = %d]\n", req.X, req.Y)
		resp := &pb.NumberRep{Result: req.X + req.Y}
		stream.Send(resp)
		fmt.Printf("Server Sended: [Result = %d]\n", resp.Result)
	}
	return nil
}

func main() {
	fmt.Printf("启动服务[Port=%d]...\n", Port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterCounterServer(s, &server{})
	s.Serve(conn)
}
