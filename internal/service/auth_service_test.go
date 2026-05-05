package service

import (
	"testing"

	"go-chat-app/internal/model"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepo struct{}

func (m *MockUserRepo) Create(user *model.User) error {
	return nil
}

func (m *MockUserRepo) FindByEmail(email string) (*model.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	return &model.User{
		ID:       1,
		Email:    email,
		Password: string(hash),
	}, nil
}

func TestRegister(t *testing.T) {
	repo := &MockUserRepo{}
	service := &AuthService{UserRepo: repo}

	req := &model.RegisterRequest{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "123456",
	}

	err := service.Register(req)

	assert.Nil(t, err)
}

func TestLoginSuccess(t *testing.T) {
	repo := &MockUserRepo{}
	service := &AuthService{UserRepo: repo}

	req := &model.LoginRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	token, err := service.Login(req)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWrongPassword(t *testing.T) {
	repo := &MockUserRepo{}
	service := &AuthService{UserRepo: repo}

	req := &model.LoginRequest{
		Email:    "test@mail.com",
		Password: "wrong",
	}

	token, err := service.Login(req)

	assert.NotNil(t, err)
	assert.Empty(t, token)
}
