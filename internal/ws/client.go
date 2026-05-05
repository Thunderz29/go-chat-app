package ws

import (
	"encoding/json"
	"go-chat-app/internal/model"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	userID int
}

type Message struct {
	To      int    `json:"to"`
	Message string `json:"message"`
}

func (c *Client) Read(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		log.Println("RAW MESSAGE:", string(msg))

		var m Message
		if err := json.Unmarshal(msg, &m); err != nil {
			log.Println("JSON ERROR:", err)
			continue
		}

		// VALIDASI
		if c.userID == 0 || m.To == 0 {
			log.Println("INVALID USER ID:", c.userID, "->", m.To)
			continue
		}

		log.Println("SAVE MESSAGE:", c.userID, "->", m.To, m.Message)

		err = hub.messageRepo.Create(&model.Message{
			SenderID:   c.userID,
			ReceiverID: m.To,
			Content:    m.Message,
		})

		if err != nil {
			log.Println("DB ERROR:", err)
		}

		target, ok := hub.clients[m.To]
		if ok {
			target.send <- []byte(m.Message)
		} else {
			log.Println("TARGET NOT CONNECTED:", m.To)
		}
	}
}

func (c *Client) Write() {
	defer c.conn.Close()

	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
