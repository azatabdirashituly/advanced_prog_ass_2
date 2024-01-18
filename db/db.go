// db/db.go
package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func DbConnection() error {
	clientOptions := options.Client().ApplyURI("mongodb+srv://armornurik:Qwerty31537597@cluster.3qjtl56.mongodb.net/")
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	CreateIndexes()

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")

	// Вернуть nil, если соединение успешно
	return nil
}
