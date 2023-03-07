package main

import (
	"fmt"
	pb "golang-rpc-stream-demo/proto/stream_demo"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct{}

func (s Server) OpenAiChat(req *pb.Request, srv pb.StreamService_OpenAiChatServer) error {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()
			time.Sleep(time.Duration(count) * time.Second)
			resp := pb.Response{Result: fmt.Sprintf("Request #%d For Id:%v", count, req.Input)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}
	
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, Server{})
	
	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
