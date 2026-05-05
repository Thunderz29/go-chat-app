package main

import (
	"log"
	"net/http"

	"go-chat-app/internal/config"
	"go-chat-app/internal/handler"
	"go-chat-app/internal/middleware"
	"go-chat-app/internal/repository"
	"go-chat-app/internal/service"
	"go-chat-app/internal/ws"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB()

	// 🟢 INIT REPO
	userRepo := &repository.UserRepository{}
	msgRepo := &repository.MessageRepository{}

	// 🟢 INIT WS HUB (HARUS DI ATAS)
	hub := ws.NewHub(msgRepo)
	go hub.Run()

	// 🟢 INIT SERVICE
	authService := &service.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{Service: authService}

	// 🟢 HANDLER YANG BUTUH HUB
	userHandler := &handler.UserHandler{Hub: hub}
	msgHandler := &handler.MessageHandler{Repo: msgRepo}

	// Routes
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", middleware.JWTMiddleware(handler.Profile))
	http.HandleFunc("/online-users", userHandler.GetOnlineUsers)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(hub, w, r)
	})

	http.HandleFunc("/messages", msgHandler.GetConversation)

	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
