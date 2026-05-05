package handler

import (
	"encoding/json"
	"net/http"

	"go-chat-app/internal/ws"
)

type UserHandler struct {
	Hub *ws.Hub
}

func (h *UserHandler) GetOnlineUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Hub.GetOnlineUserIDs()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
