package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/luxgru35/todo-app"
	"github.com/luxgru35/todo-app/pkg/repository"
)

const salt = "asdsadfsaghhboopbvolffh3423423gh"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return s.repo.CreateUser(user)
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
