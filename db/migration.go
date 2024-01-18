package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes() error {
	// Установка опций клиента
	clientOptions := options.Client().ApplyURI("mongodb+srv://armornurik:Qwerty31537597@cluster.3qjtl56.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Создание индексов для коллекции "volunteers"
	studentCollection := client.Database("Learn").Collection("students")
	_, err = studentCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    map[string]interface{}{"phone": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}

	// Создание индексов для коллекции "children"
	teacherCollection := client.Database("Learn").Collection("teachers")
	_, err = teacherCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    map[string]interface{}{"phone": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
