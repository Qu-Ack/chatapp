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

func handleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error while upgrading the socket connection")
		return
	}

	Read(conn)

}

func Read(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("message recieved: %s", message)

		// let the client know his message was recieved

		err = conn.WriteMessage(websocket.TextMessage, []byte("success"))
		if err != nil {
			log.Println("error while writing the connect status to client")
			break
		}
	}
}
