package config

import (
	"boilerplate/exception"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Database {
	mongoURL := os.Getenv("MONGO_HOST")

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoURL)

	if os.Getenv("MONGO_LOG_QUERY") == "1" {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				fmt.Println(evt.Command)
			},
		}
		clientOptions.SetMonitor(cmdMonitor)
	}

	client, err := mongo.NewClient(clientOptions)
	exception.PanicIfNeeded(err)

	err = client.Connect(context.Background())
	exception.PanicIfNeeded(err)

	dbName := os.Getenv("MONGO_NAME")
	return client.Database(dbName)
}
