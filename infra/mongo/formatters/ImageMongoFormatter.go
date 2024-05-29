package formatters

import (
	"GoDockerSandbox/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

type ImageMongoFormatter struct {
}

func NewImageMongoFormatter() *ImageMongoFormatter {
	return &ImageMongoFormatter{}
}

func (fmtr *ImageMongoFormatter) FormatImage(image model.Image) interface{} {
	doc := bson.M{
		"id":        image.Id,
		"ImageName": image.ImageName,
		"name":      image.Name,
		"tag":       image.Tag,
		"isInfra":   image.IsInfra,
		"ports":     image.Ports,
		"envs":      image.Envs,
	}
	return doc
}

func (fmtr *ImageMongoFormatter) FormatImages(images []model.Image) []interface{} {
	docs := make([]interface{}, 0, len(images))
	for _, image := range images {
		docs = append(docs, fmtr.FormatImage(image))
	}
	return docs
}
