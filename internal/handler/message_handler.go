package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-chat-app/internal/repository"
)

type MessageHandler struct {
	Repo *repository.MessageRepository
}

func (h *MessageHandler) GetConversation(w http.ResponseWriter, r *http.Request) {
	user1Str := r.URL.Query().Get("user1")
	user2Str := r.URL.Query().Get("user2")

	user1, _ := strconv.Atoi(user1Str)
	user2, _ := strconv.Atoi(user2Str)

	messages, err := h.Repo.GetConversation(user1, user2)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
