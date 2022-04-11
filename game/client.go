// Handle sportsdata.io API polling requests and push updates to WebSocket event hub.

package game

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sb-server/stream"
	"sort"
	"strings"
	"time"
)

const BASE_URL = "https://api.sportsdata.io/v3/nba/scores/json/"
const GAMES_BY_DATE_URL = BASE_URL + "GamesByDate"
const API_KEY = "b3ba3ff994b54af28a8e850b5ee7a419"

func GetGamesByDate(date time.Time) ([]Game, error) {
	resp, err := http.Get(GAMES_BY_DATE_URL +
		strings.ToUpper(date.Format("/2006-Jan-02")) +
		"?key=" + API_KEY)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	//sb := string(body)
	if err != nil {
		return nil, err
	}

	var games []Game
	json.Unmarshal(body, &games)

	sort.Slice(games, func(i, j int) bool { return games[i].DateTime.Time.Before(games[j].DateTime.Time) })

	return games, nil
}

func PushGameUpdates(h *stream.Hub) {
	for {
		games, err := GetGamesByDate(time.Now())
		<-time.After(10 * time.Hour)
		if err != nil {
			print(err)
		}
		h.Broadcast <- stream.SocketMessage{MessageType: stream.GameUpdates, Data: &games}
	}
}
