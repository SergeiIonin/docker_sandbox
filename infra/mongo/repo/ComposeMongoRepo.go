package repo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"GoDockerSandbox/domain/model"
	"GoDockerSandbox/infra/mongo/formatters"
)

type ComposeMongoRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	formatter  *formatters.ComposeMongoFormatter
}

func NewComposeMongoRepo(uri string) (*ComposeMongoRepo, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	database := client.Database("docker_sandbox")
	collection := database.Collection("composes")
	fmtr := formatters.NewComposeMongoFormatter()

	return &ComposeMongoRepo{
		client:     client,
		database:   database,
		collection: collection,
		formatter:  fmtr,
	}, nil
}

func (r *ComposeMongoRepo) Get(id string) (model.Compose, error) {
	filter := bson.M{"id": id}
	var compose model.Compose
	err := r.collection.FindOne(context.Background(), filter).Decode(&compose)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Compose{}, fmt.Errorf("compose not found")
		}
		return model.Compose{}, err
	}
	return compose, nil
}

func (r *ComposeMongoRepo) GetAll() ([]model.Compose, error) {
	filter := bson.D{}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("composes not found")
		}
		return nil, err
	}
	composes := []model.Compose{}

	for cursor.Next(context.Background()) {
		var compose model.Compose
		err := cursor.Decode(&compose)
		if err != nil {
			return nil, err
		}
		composes = append(composes, compose)
	}

	return composes, nil
}

func (r *ComposeMongoRepo) Save(compose model.Compose) (id string, err error) {
	doc := r.formatter.FormatCompose(compose)
	_, err = r.collection.InsertOne(context.Background(), doc)
	id = compose.Id
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *ComposeMongoRepo) Upsert(compose model.Compose) (id string, err error) {
	doc := r.formatter.FormatCompose(compose)
	id = compose.Id
	findDoc := bson.M{"id": id}
	res := r.collection.FindOneAndReplace(context.Background(), findDoc, doc)
	if res.Err() != nil {
		return id, res.Err()
	}
	return
}

func (r *ComposeMongoRepo) Update(id string, yaml string) (string, error) {
	doc := bson.M{"$set": bson.M{"yaml": yaml}}
	filter := bson.M{"id": id}
	_, err := r.collection.UpdateOne(context.Background(), filter, doc)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *ComposeMongoRepo) Delete(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *ComposeMongoRepo) DeleteAll() error {
	_, err := r.collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
