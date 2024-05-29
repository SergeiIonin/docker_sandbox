package formatters

import (
	"GoDockerSandbox/domain/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ComposeMongoFormatter struct {
}

func NewComposeMongoFormatter() *ComposeMongoFormatter {
	return &ComposeMongoFormatter{}
}

func (fmtr *ComposeMongoFormatter) FormatCompose(compose model.Compose) interface{} {
	doc := bson.M{
		"id":          compose.Id,
		"name":        compose.Name,
		"appImages":   compose.AppImages,
		"infraImages": compose.InfraImages,
		"network":     compose.Network,
		"yaml":        compose.Yaml,
	}
	return doc
}

func (fmtr *ComposeMongoFormatter) FormatComposes(composes []model.Compose) []interface{} {
	docs := make([]interface{}, 0, len(composes))
	for _, compose := range composes {
		docs = append(docs, fmtr.FormatCompose(compose))
	}
	return docs
}
