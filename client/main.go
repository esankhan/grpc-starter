package main

import (
	"log"

	pb "github.com/esankhan/portfolio/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)


func main() {
	conn,err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn);

	 names := &pb.NamesList {
		Names: []string{"Ehsan","Ali","Sankhan"},
	 }	
	CallSayHelloServerStream(client,names)
	//CallSayHello(client)
}