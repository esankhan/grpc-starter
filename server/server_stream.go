package main

import (
	"log"
	"time"

	pb "github.com/esankhan/portfolio/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList,stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		log.Printf("Sent: %v", res)
		stream.Send(&pb.HelloResponse{
			Message: "Hello " + name,
		})

		time.Sleep(1 * time.Second)
	}
	return nil
}