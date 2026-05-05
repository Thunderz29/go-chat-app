package repository

import (
	"go-chat-app/internal/config"
	"go-chat-app/internal/model"
)

type MessageRepository struct{}

func (r *MessageRepository) Create(msg *model.Message) error {
	query := `
	INSERT INTO messages (sender_id, receiver_id, content)
	VALUES (?, ?, ?)
	`
	_, err := config.DB.Exec(query, msg.SenderID, msg.ReceiverID, msg.Content)
	return err
}

func (r *MessageRepository) GetConversation(user1, user2 int) ([]model.Message, error) {
	var messages []model.Message

	query := `
	SELECT * FROM messages
	WHERE (sender_id = ? AND receiver_id = ?)
	   OR (sender_id = ? AND receiver_id = ?)
	ORDER BY created_at ASC
	`

	err := config.DB.Select(&messages, query, user1, user2, user2, user1)
	return messages, err
}
