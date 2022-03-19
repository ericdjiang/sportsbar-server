// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sb-server/game"
	"sb-server/stream"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveSports(w http.ResponseWriter, r *http.Request) {
	games, err := game.GetGamesByDate()
	if err != nil {

	}
	fmt.Println(games)
	fmt.Fprintf(w, "done")
}

func main() {
	flag.Parse()
	hub := stream.NewHub()
	go hub.Run()
	go game.PushGameUpdates(hub)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		stream.ServeWs(hub, w, r)
	})
	http.HandleFunc("/sports", serveSports)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
