package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"context"
	"github.com/sigurniv/grpc-web-golang/service/steamgame/steamgamepb"
)

func main() {
	fmt.Println("Hello from client! :)")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := steamgamepb.NewSteamGameServiceClient(cc)

	response := doGameSearch(c)
	if (len(response.Games) > 0) {
		doGameDetails(c, response.Games[0].Id)
	}

}

func doGameSearch(c steamgamepb.SteamGameServiceClient) *steamgamepb.SearchGameResponse {
	fmt.Println("Starting to do Game Search Request RPC...")

	req := &steamgamepb.SearchGameRequest{Search: &steamgamepb.GameSearch{Search: "lol"}}
	res, err := c.SearchGame(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Game Search RPC: %v", err)
	}

	log.Printf("Response from Game Search: %v", res.Games)

	return res
}

func doGameDetails(c steamgamepb.SteamGameServiceClient, id string) {
	log.Printf("Starting to do Game Details Request for id: %v", id)

	req := &steamgamepb.GameDetailsRequest{Game: &steamgamepb.Game{Id: id}}
	res, err := c.GameDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Game Search RPC: %v", err)
	}

	log.Printf("Response from Game Details: %v", res.Details)
}
