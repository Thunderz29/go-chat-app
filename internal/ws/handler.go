package ws

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil || userID == 0 {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	log.Println("CONNECT USER:", userID)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("UPGRADE ERROR:", err)
		return
	}

	log.Println("UPGRADE SUCCESS:", userID)

	client := &Client{
		conn:   conn,
		send:   make(chan []byte),
		userID: userID,
	}

	log.Println("BEFORE REGISTER:", userID)
	hub.register <- client
	log.Println("AFTER REGISTER:", userID)

	go client.Read(hub)
	go client.Write()
}
