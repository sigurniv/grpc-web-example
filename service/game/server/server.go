package server

import (
	"github.com/sigurniv/grpc-web-golang/service/game/gamepb"
	"context"
	"log"
	"github.com/sigurniv/grpc-web-golang/service/game/steamgamepb"
	"google.golang.org/grpc"
)

type server struct {
}

func New() *server {
	return &server{}
}

const host = "steamgame.api:50051";

func (*server) SearchGame(ctx context.Context, req *gamepb.SearchGameRequest) (*gamepb.SearchGameResponse, error) {
	cc, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := steamgamepb.NewSteamGameServiceClient(cc)

	rq := &steamgamepb.SearchGameRequest{Search: &steamgamepb.GameSearch{Search: req.Search.Search}}
	res, err := c.SearchGame(context.Background(), rq)
	if err != nil {
		log.Fatalf("error while calling Game Search RPC: %v", err)
	}

	log.Printf("Response from Steam Game Search: %v", res.Games)

	var respGames []*gamepb.Game
	for _, game := range res.Games {
		respGames = append(respGames, &gamepb.Game{
			Id:    game.Id,
			Title: game.Title,
		})
	}
	return &gamepb.SearchGameResponse{Games: respGames}, nil
}

func (*server) GameDetails(ctx context.Context, req *gamepb.GameDetailsRequest) (*gamepb.GameDetailsResponse, error) {
	cc, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := steamgamepb.NewSteamGameServiceClient(cc)

	log.Println("Starting Steam Game Details request...")
	rq := &steamgamepb.GameDetailsRequest{Game: &steamgamepb.Game{Id: req.Game.Id}}
	res, err := c.GameDetails(context.Background(), rq)
	if err != nil {
		log.Fatalf("error while calling Game Search RPC: %v", err)
	}

	log.Printf("Response from Steam Game Details: %v", res.Details)

	return &gamepb.GameDetailsResponse{Details: &gamepb.GameDetails{
		Id:       res.Details.Id,
		Price:    res.Details.Price,
		Name:     res.Details.Name,
		Currency: res.Details.Currency,
	}}, nil
}
