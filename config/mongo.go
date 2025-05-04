package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB не отвечает:", err)
	}

	fmt.Println("✅ Подключение к MongoDB успешно")
	Client = client
}
