// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// https://github.com/gorilla/websocket/blob/master/LICENSE

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	hubManager := newHubManager()
	flag.Parse()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		print("/ws accessed")
		serveWs(hubManager, w, r)
	})

	hubManager_matching := newHubManager()
	http.HandleFunc("/ws_matching", func(w http.ResponseWriter, r *http.Request) {
		serveWs_matching(hubManager_matching, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
