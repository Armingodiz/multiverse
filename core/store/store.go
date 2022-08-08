package store

import (
	"context"
	"errors"
	"multiverse/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoStore(collection *mongo.Collection) Store {
	return &MongoStore{Collection: collection}
}

type Store interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, email string) (*models.User, error)
	DeleteUser(ctx context.Context, email string) error
}
type MongoStore struct {
	Collection *mongo.Collection
}

func (s *MongoStore) CreateUser(ctx context.Context, user *models.User) error {
	res, err := s.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("cannot convert to OID")
	}
	return nil
}

func (s *MongoStore) GetUser(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	filter := bson.M{"email": email}
	res := s.Collection.FindOne(ctx, filter)
	if err := res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *MongoStore) DeleteUser(ctx context.Context, email string) error {
	filter := bson.M{"email": email}
	res, err := s.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
