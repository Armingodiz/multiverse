package store

import (
	"context"
	"multiverse/core/models"
)

func NewStore() Store {
	return &store{}
}

type Store interface {
	CreateUser(ctx context.Context, user *models.User) error
	//todo:add needed methods for your database
}
type store struct {
}

func (s *store) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}
