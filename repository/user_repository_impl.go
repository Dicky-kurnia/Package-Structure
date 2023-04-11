package repository

import (
	"boilerplate/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var konteks = context.Background()

type userRepository struct {
	Mongo *mongo.Database
}

func (repository userRepository) GetUser(username string) (entity.UserEntity, error) {
	userCollection := repository.Mongo.Collection(entity.UserEntity{}.CollectionName())

	result := entity.UserEntity{}
	err := userCollection.FindOne(konteks, bson.M{"username": username}).Decode(&result)
	return result, err
}

func NewUserRepository(mongo *mongo.Database) UserRepository {
	return &userRepository{Mongo: mongo}
}
