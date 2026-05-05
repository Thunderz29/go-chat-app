package service

import (
	"errors"
	"os"
	"time"

	"go-chat-app/internal/model"
	"go-chat-app/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repository.UserRepositoryInterface
}

func (s *AuthService) Register(req *model.RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.UserRepo.Create(&user)
}

func (s *AuthService) Login(req *model.LoginRequest) (string, error) {
	user, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
