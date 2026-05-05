# 🚀 Go Realtime Chat App (7 Days Project)

## 📌 Overview

Project ini adalah backend **Realtime Chat System** menggunakan Golang yang dibangun dalam 7 hari.

Fitur utama:

* 🔐 JWT Authentication
* 💬 Realtime Chat (WebSocket)
* 💾 Message Persistence (MySQL)
* 📡 Online/Offline Presence
* 📜 Chat History API
* 🧪 Unit Testing (Service Layer)

---

# 🧱 Tech Stack

* Golang (net/http)
* MySQL
* Gorilla WebSocket
* JWT (golang-jwt)
* Bcrypt
* Testify (unit testing)

---

# 🏗️ Project Structure

```
go-chat-app/
│── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── middleware/
│   └── ws/
│
├── .env
├── go.mod
└── README.md
```

---

# ⚙️ Setup & Run

## 1. Clone project

```
git clone <repo-url>
cd go-chat-app
```

## 2. Install dependency

```
go mod tidy
```

## 3. Setup MySQL

```
CREATE DATABASE chat_app;
```

## 4. Create Tables

### users

```
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### messages

```
CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sender_id INT,
    receiver_id INT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id)
);
```

## 5. Setup `.env`

```
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=chat_app

JWT_SECRET=supersecretkey
```

## 6. Run app

```
go run cmd/main.go
```

---

# 🔑 API Documentation

## 🟢 Auth

### Register

```
POST /register
```

Body:

```json
{
  "name": "user1",
  "email": "user1@mail.com",
  "password": "123456"
}
```

---

### Login

```
POST /login
```

Response:

```
JWT TOKEN
```

---

### Profile (Protected)

```
GET /profile
Authorization: Bearer <token>
```

---

## 💬 Chat

### WebSocket

```
ws://localhost:8080/ws?user_id=1
```

### Send Message

```json
{
  "to": 2,
  "message": "halo"
}
```

---

## 📜 Chat History

```
GET /messages?user1=1&user2=2
```

---

## 📡 Online Users

```
GET /online-users
```

---

# 🧪 Testing

## Unit Test

```
go test ./...
```

---

## Manual Testing

### 1. Connect WebSocket

```javascript
let ws1 = new WebSocket("ws://localhost:8080/ws?user_id=1")
let ws2 = new WebSocket("ws://localhost:8080/ws?user_id=2")
```

### 2. Send Message

```javascript
ws1.send(JSON.stringify({
  to: 2,
  message: "halo"
}))
```

---

# 📅 Development Timeline (7 Days)

## Day 1 — Setup

* Init project
* Setup MySQL
* DB connection
* Schema design

---

## Day 2 — Authentication

* Register user
* Hash password (bcrypt)
* Login
* JWT token
* Middleware auth

---

## Day 3 — WebSocket Basic

* Setup WebSocket endpoint
* Handle connection
* Hub (connection manager)

---

## Day 4 — Realtime Chat

* Send message antar user
* Routing message
* Concurrency (goroutine + channel)

---

## Day 5 — Persistence

* Save message ke MySQL
* Get chat history
* API endpoint history

---

## Day 6 — Presence System

* Online/offline tracking
* Broadcast status
* Endpoint online users

---

## Day 7 — Final

* Unit testing
* Refactor clean code
* Logging improvement
* Documentation

---

# 🧠 Key Concepts

* WebSocket realtime communication
* Goroutine & channel (concurrency)
* Clean architecture (handler-service-repo)
* JWT authentication flow
* Database relationship (foreign key)
* Dependency Injection via interface
* Unit testing with mock

---

# 🎯 Learning Outcome

Setelah project ini, kamu memahami:

* Cara membangun backend realtime system
* Concurrency di Golang
* Best practice struktur backend
* Cara membuat project siap interview

---

# 🔥 Future Improvements

* Redis (scaling presence)
* Kafka (event-driven)
* Group chat
* Read receipt
* Docker & deployment
* Rate limiting

---

# 💡 Author Note

Project ini dibuat sebagai latihan backend engineering dalam 7 hari, dengan fokus pada:

* real-world use case
* clean architecture
* scalable design

---

# ⭐ Conclusion

Ini bukan sekadar CRUD project, tapi:

👉 **Realtime Chat System dengan concurrency, persistence, dan authentication**

---

# 🚀 Ready for Interview

Kalau kamu bisa menjelaskan project ini dengan baik, kamu sudah berada di level:

👉 **Junior → Mid Backend Engineer**
