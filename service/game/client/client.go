package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"context"
	"github.com/sigurniv/grpc-web-golang/service/game/gamepb"
)

func main() {
	fmt.Println("Hello from client! :)")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := gamepb.NewGameServiceClient(cc)

	response := doGameSearch(c)

	if len(response.Games) > 0 {
		doGameDetails(c, response.Games[0].Id)
	}

}

func doGameSearch(c gamepb.GameServiceClient) *gamepb.SearchGameResponse {
	fmt.Println("Starting to do Game Search Request RPC...")

	req := &gamepb.SearchGameRequest{Search: &gamepb.GameSearch{Search: "lol"}}
	res, err := c.SearchGame(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Game Search RPC: %v", err)
	}

	log.Printf("Response from Game Search: %v", res.Games)

	return res
}

func doGameDetails(c gamepb.GameServiceClient, id string) {
	log.Printf("Starting to do Game Details Request for id: %v", id)

	req := &gamepb.GameDetailsRequest{Game: &gamepb.Game{Id: id}}
	res, err := c.GameDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Steam Game Details RPC: %v", err)
	}

	log.Printf("Response from Game Details: %v", res.Details)
}
