package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket Endpoint hit")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error while upgrading the socket connection")
		return
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}
