package repository

import (
	"go-chat-app/internal/config"
	"go-chat-app/internal/model"
)

type UserRepositoryInterface interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type UserRepository struct{}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM users WHERE email = ?`
	err := config.DB.Get(&user, query, email)

	return &user, err
}
