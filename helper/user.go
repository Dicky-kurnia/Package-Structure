package helper

import (
	"boilerplate/entity"
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func InsertUser(mongo *mongo.Database) {
	username := StringPrompt("Username:")
	password := StringPrompt("Password:")
	userCollection := mongo.Collection(entity.UserEntity{}.CollectionName())

	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	_, err = userCollection.InsertOne(context.Background(), entity.UserEntity{
		Id:       primitive.NewObjectID(),
		Username: username,
		Password: string(pass),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucess!")
}
