package rawg_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GameResult struct {
	ID              int     `json:"id"`
	Slug            string  `json:"slug"`
	Description     string  `json:"description"`
	Metacritic      int     `json:"metacritic"`
	MetacriticUrl   string  `json:"metacritic_url"`
	BackgroundImage string  `json:"background_image"`
	Publisher       string  `json:"publisher"`
	Rating          float32 `json:"rating"`
}

var RawgApiEndpoint = "https://api.rawg.io/api"

var gameResult GameResult

func GetGameDetails(rawgId string) GameResult {
	path := fmt.Sprintf("games/%s", rawgId)
	return doRequest(path)
}

func doRequest(path string) GameResult {
	RawgApiKey := os.Getenv("RAWG_API_KEY")
	url := fmt.Sprintf("%s/%s?key=%s", RawgApiEndpoint, path, RawgApiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		errorMessage := fmt.Sprintf("Rawg request Error - Path: %s : %s ", path, resp.Status)
		log.Println(errorMessage)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GameResult{}
	}
	err = json.Unmarshal(body, &gameResult)
	if err != nil {
		log.Println(err.Error())
	}

	return gameResult

}
