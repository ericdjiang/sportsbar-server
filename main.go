// HTTP Server entry point which serves a websocket connection URL and RESTful GET endpoint.

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"sb-server/game"
	"sb-server/stream"
	"strconv"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveGamesPrevNDays(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["days"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		http.Error(w, "The required query parameter 'days' is missing.", http.StatusBadRequest)
		return
	}
	nDays, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Please ensure 'days' is an integer parameter.", http.StatusBadRequest)
		return
	}

	date := time.Now()
	var games []game.Game
	for i := 1; i <= nDays; i++ {
		gamesByDate, err := game.GetGamesByDate(date)
		if err != nil {
			http.Error(w, "Error contacting SportsData API.", http.StatusBadRequest)
			return
		}
		games = append(games, gamesByDate...)
		date = date.AddDate(0, 0, -i)
	}
	json.NewEncoder(w).Encode(games)
}

func main() {
	flag.Parse()
	hub := stream.NewHub()
	go hub.Run()
	go game.PushGameUpdates(hub)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		stream.ServeWs(hub, w, r)
	})
	http.HandleFunc("/games", serveGamesPrevNDays)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Setting port to 8080")
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
