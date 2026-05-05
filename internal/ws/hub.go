package ws

import (
	"encoding/json"
	"go-chat-app/internal/repository"
	"log"
)

type Hub struct {
	clients    map[int]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte

	messageRepo *repository.MessageRepository
}

func NewHub(repo *repository.MessageRepository) *Hub {
	return &Hub{
		clients:     make(map[int]*Client),
		register:    make(chan *Client, 16),
		unregister:  make(chan *Client, 16),
		broadcast:   make(chan []byte, 32),
		messageRepo: repo,
	}
}

func (h *Hub) notifyStatus(userID int, status string) {
	data := map[string]interface{}{
		"type":   "status",
		"userId": userID,
		"status": status,
	}

	msg, _ := json.Marshal(data)

	select {
	case h.broadcast <- msg:
	default:
		log.Println("[WS] broadcast full")
	}
}

func (h *Hub) Run() {
	log.Println("[WS] HUB RUNNING")

	for {
		select {
		case client := <-h.register:
			log.Println("[WS] REGISTER:", client.userID)

			h.clients[client.userID] = client
			h.notifyStatus(client.userID, "online")

		case client := <-h.unregister:
			log.Println("[WS] UNREGISTER:", client.userID)

			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}

			h.notifyStatus(client.userID, "offline")

		case message := <-h.broadcast:
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					log.Println("[WS] drop message to", client.userID)
				}
			}
		}
	}
}

func (h *Hub) GetOnlineUserIDs() []int {
	var ids []int
	for id := range h.clients {
		ids = append(ids, id)
	}
	return ids
}
