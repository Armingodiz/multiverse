package userService

import (
	"context"
	"multiverse/core/models"
	"multiverse/core/store"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
	DeleteUser(email string) error
}

var simpleContext = context.Background()

func NewUserService(store store.Store) UserService {
	return &userService{
		Store: store,
	}
}

type userService struct {
	Store store.Store
}

func (s *userService) CreateUser(user *models.User) error {
	return s.Store.CreateUser(simpleContext, user)
}

func (s *userService) GetUser(email string) (*models.User, error) {
	return s.Store.GetUser(simpleContext, email)
}

func (s *userService) DeleteUser(email string) error {
	return s.Store.DeleteUser(simpleContext, email)
}
