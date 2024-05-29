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

type ImageMongoRepo struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	formatter  *formatters.ImageMongoFormatter
}

func NewImageMongoRepo() (*ImageMongoRepo, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
	collection := database.Collection("app_images")
	fmtr := formatters.NewImageMongoFormatter()

	return &ImageMongoRepo{
		client:     client,
		database:   database,
		collection: collection,
		formatter:  fmtr,
	}, nil
}

func (r *ImageMongoRepo) Get(id string) (model.Image, error) {
	filter := bson.M{"id": id}
	var image model.Image
	err := r.collection.FindOne(context.Background(), filter).Decode(&image)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Image{}, fmt.Errorf("image not found")
		}
		return model.Image{}, err
	}
	return image, nil
}

func (r *ImageMongoRepo) GetAll() ([]model.Image, error) {
	filter := bson.D{}
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("images not found")
		}
		return nil, err
	}
	images := []model.Image{}

	for cursor.Next(context.Background()) {
		var image model.Image
		err := cursor.Decode(&image)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return images, nil
}

func (r *ImageMongoRepo) Save(image model.Image) error {
	doc := r.formatter.FormatImage(image)
	_, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *ImageMongoRepo) SaveAll(images []model.Image) error {
	docs := r.formatter.FormatImages(images)
	_, err := r.collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}
	return nil
}

func (r *ImageMongoRepo) Update(image model.Image) error {
	doc := r.formatter.FormatImage(image)
	filter := bson.M{"id": image.Id}
	_, err := r.collection.ReplaceOne(context.Background(), filter, doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *ImageMongoRepo) UpdateAll(images []model.Image) error {
	docs := r.formatter.FormatImages(images)
	for _, doc := range docs {
		filter := bson.M{"id": doc.(bson.M)["id"]}
		_, err := r.collection.ReplaceOne(context.Background(), filter, doc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ImageMongoRepo) Delete(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *ImageMongoRepo) DeleteAll() error {
	filter := bson.D{}
	_, err := r.collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
