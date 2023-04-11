package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}

func (UserEntity) CollectionName() string {
	return "users"
}
