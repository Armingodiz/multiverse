package userService

import (
	"multiverse/core/models"
	"multiverse/core/store"
)

type UserService interface {
	CreateUser(user *models.User) error
}

func NewUserService(store *store.Store) UserService {
	return &userService{
		Store: store,
	}
}

type userService struct {
	Store *store.Store
}

func (s *userService) CreateUser(user *models.User) error {
	return nil
}
