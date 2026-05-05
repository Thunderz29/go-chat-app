package model

import "time"

type Message struct {
	ID         int       `db:"id"`
	SenderID   int       `db:"sender_id"`
	ReceiverID int       `db:"receiver_id"`
	Content    string    `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
}
