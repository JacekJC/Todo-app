package main

import(
	"fmt"
)

type ConnManager struct{
	clients map[*Client]bool

	names map[string]*Client

	register chan *Client

	unregister chan*Client

	broadcast chan []byte
}

func new_connmanager() *ConnManager {
	fmt.Println("new manager made");
	return &ConnManager{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make (map[*Client]bool),
		names: make (map[string]*Client),
	}
}

func (h *ConnManager) run() {
	for{
		select{
			case client := <-h.register:
				h.clients[client] = true
			case client := <-h.unregister:
				if _, ok := h.clients[client]; ok{
					delete(h.clients, client)
					close(client.send)
				}
			case message := <-h.broadcast:
				for client := range h.clients{
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
				}
			}
		}
	}
}

