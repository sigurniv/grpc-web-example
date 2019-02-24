package main

import (
	"log"
	"fmt"
	"net"
	"google.golang.org/grpc"
	"github.com/sigurniv/grpc-web-golang/service/steamgame/steamgamepb"
	"os"
	"os/signal"
	"github.com/sigurniv/grpc-web-golang/service/steamgame/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	srv := server.New()

	fmt.Println("Steam Game Steam Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	steamgamepb.RegisterSteamGameServiceServer(s, srv)

	go func() {
		log.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Println("Stopping the server")
	s.Stop()
	log.Println("Closing the listener")
	lis.Close()
	log.Println("End of Program")
}
