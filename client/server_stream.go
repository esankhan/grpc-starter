package main

import (
	"context"
	"io"
	"log"

	pb "github.com/esankhan/portfolio/proto"
)


func CallSayHelloServerStream(client pb.GreetServiceClient,names *pb.NamesList) {
	log.Printf("Server streaming started");
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Error while calling SayHelloServerStreaming: %v", err)
	}
	for {
		message,err := stream.Recv();
		
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}

		log.Printf("Received: %v", message)
	}

	log.Printf("Server streaming ended");
}