package services

import (
	"context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

type InfraImageRepo struct {
	client *mongo.Client
}

func NewInfraImageRepo(mongoUri string) *InfraImageRepo {
    clientOptions := options.Client().ApplyURI(mongoUri) // "mongodb://localhost:27017"
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    return &InfraImageRepo{client: client}
}

