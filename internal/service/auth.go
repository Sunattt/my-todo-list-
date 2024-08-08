package service

import (
	"crypto/sha1"
	"fmt"
	repository "todo/internal/repository/psgdb"
	"todo/pkg/todo_list"
)
const salt = "bnapb1@#$%^&wvk425q3tebnjoh[ipqn]"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService (repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_list.User) (int, error){
	user.Password = s.generationPassword(user.Password) 
	return s.repo.CreateUser(user)
}

func(s *AuthService) generationPassword (password string ) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}