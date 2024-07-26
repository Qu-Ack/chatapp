package main

import "log"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func newPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}

}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			log.Println("Size of connection pool : ", len(p.Clients))
			for client, _ := range p.Clients {
				log.Println(client)
				client.Conn.WriteJSON(Message{Type: 3, Body: "New User Joined ...."})
			}
			break

		case client := <-p.Unregister:
			delete(p.Clients, client)
			log.Println("Size of connection pool : ", len(p.Clients))
			for client, _ := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 3, Body: "A User left ...."})
			}
			break
		case message := <-p.Broadcast:
			log.Println("sending message to all clients")
			for client, _ := range p.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Println(err)
					return
				}
			}

		}
	}
}
