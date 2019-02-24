package server

import (
	"log"
	"github.com/sigurniv/grpc-web-golang/service/steamgame/steamgamepb"
	"context"
	"github.com/sigurniv/grpc-web-golang/service/steamgame/steam"
	"strconv"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (*Server) SearchGame(ctx context.Context, req *steamgamepb.SearchGameRequest) (*steamgamepb.SearchGameResponse, error) {
	log.Printf("Search game request: %v", req)
	steamSvc := steam.New()
	steamGames, err := steamSvc.GameSearch(req.Search.Search)
	if err != nil {
		log.Printf("Error fetching games from steam: %v", err)
	}

	log.Printf("Games: %v", steamGames)

	var games []*steamgamepb.Game
	for _, game := range steamGames {
		games = append(games, &steamgamepb.Game{
			Id:    strconv.Itoa(game.AppId),
			Title: game.Name,
		})
	}

	return &steamgamepb.SearchGameResponse{Games: games}, nil
}

func (*Server) GameDetails(ctx context.Context, req *steamgamepb.GameDetailsRequest) (*steamgamepb.GameDetailsResponse, error) {
	log.Printf("Game details request: %v", req)
	steamSvc := steam.New()
	details, err := steamSvc.AppDetails(req.Game.Id)
	if err != nil {
		log.Printf("Error fetching game details from steam: %v", err)
	}

	log.Printf("Details: %v", details)

	return &steamgamepb.GameDetailsResponse{Details: &steamgamepb.GameDetails{
		Id:       strconv.Itoa(details.AppId),
		Name:     details.Name,
		Currency: details.Currency,
		Price:    details.Price,
	}}, nil
}