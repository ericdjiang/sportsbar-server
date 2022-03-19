package game

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sb-server/stream"
	"time"
)

const BASE_URL = "https://api.sportsdata.io/v3/nba/scores/json/"
const GAMES_BY_DATE_URL = BASE_URL + "GamesByDate"
const API_KEY = "b3ba3ff994b54af28a8e850b5ee7a419"

func GetGamesByDate() ([]Game, error) {
	resp, err := http.Get(GAMES_BY_DATE_URL +
		"/2022-MAR-14" +
		"?key=" + API_KEY)
	if err != nil {

	}

	body, err := ioutil.ReadAll(resp.Body)
	//sb := string(body)
	if err != nil {

	}

	var games []Game
	json.Unmarshal(body, &games)
	return games, nil
}

func PushGameUpdates(h *stream.Hub) {
	for {
		games, err := GetGamesByDate()
		if err != nil {

		}
		for _, game := range games {
			h.Broadcast <- stream.SocketMessage{Type: stream.GameUpdates, Data: &game}
		}
		<-time.After(5 * time.Second)
	}
}
