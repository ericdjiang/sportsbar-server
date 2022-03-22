// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
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
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
