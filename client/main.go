package main

import (
	"context"
	"fmt"
	pb "golang-rpc-stream-demo/proto/stream_demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	
	// dail server
	conn, err := grpc.Dial("localhost:50005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	
	// create stream
	client := pb.NewStreamServiceClient(conn)
	in := &pb.Request{Input: "123"}
	stream, err := client.OpenAiChat(context.Background(), in)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	
	fmt.Println("---------1111")
	//ctx := stream.Context()
	done := make(chan bool)
	
	go func() {
		for {
			fmt.Println("---------2222")
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			fmt.Println("---------3333")
			if err != nil {
				log.Fatalf("+++++++ can not receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()
	
	<-done
	log.Printf("finished")
}
