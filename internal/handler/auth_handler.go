package handler

import (
	"encoding/json"
	"net/http"

	"go-chat-app/internal/model"
	"go-chat-app/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := h.Service.Register(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("register success"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.Service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	w.Write([]byte(token))
}
