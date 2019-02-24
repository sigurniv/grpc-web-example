package steam

import (
	"net/http"
	"crypto/tls"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	"fmt"
	"errors"
	"log"
)

type Steam struct{}

func New() *Steam {
	return &Steam{
	}
}

type Game struct {
	Name  string `json:"name"`
	AppId int    `json:"appId"`
}

type GameDetails struct {
	Name     string  `json:"name"`
	AppId    int     `json:"appId"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

func (svc *Steam) GameSearch(gameName string) ([]Game, error) {
	var games []Game
	log.Println("Fetching steam games...")
	responseRaw, err := svc.request("http://api.steampowered.com/ISteamApps/GetAppList/v0001/")
	if err != nil {
		return games, err
	}

	var getAppListResponse GetAppListResponse
	err = json.Unmarshal(responseRaw, &getAppListResponse)
	if err != nil {
		return games, err
	}

	for _, item := range getAppListResponse.AppList.Apps.App {
		if strings.HasPrefix(strings.ToLower(item.Name), strings.ToLower(gameName)) {
			games = append(games, Game{
				Name:  item.Name,
				AppId: item.AppId,
			})
		}
	}

	return games, err
}

func (svc *Steam) AppDetails(appId string) (GameDetails, error) {
	var gameDetails GameDetails
	responseRaw, err := svc.request(fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%s", appId))
	if err != nil {
		return gameDetails, err
	}

	var appDetailsResponse AppDetailsResponse
	err = json.Unmarshal(responseRaw, &appDetailsResponse)
	if err != nil {
		return gameDetails, err
	}

	details, ok := appDetailsResponse[appId]
	if !ok {
		return gameDetails, errors.New("Bad response from steam")
	}

	gameDetails.Name = details.Data.Name
	gameDetails.AppId = details.Data.AppId
	gameDetails.Currency = details.Data.PriceOverview.Currency
	gameDetails.Price = details.Data.PriceOverview.Price / 100

	return gameDetails, err
}

func (svc *Steam) request(url string) ([]byte, error) {
	var response []byte
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	client.Timeout = 5 * time.Second
	respRaw, err := client.Get(url)
	if err != nil {
		return response, err
	}

	defer respRaw.Body.Close()

	return ioutil.ReadAll(respRaw.Body)
}
