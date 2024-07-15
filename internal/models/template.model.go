package models

import (
	"github.com/sreerag-rajan/boilerplate-go/config"
	modelImplementation "github.com/sreerag-rajan/boilerplate-go/pkg/database/model_implementation"
)

type Template struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Archived string `json:"archived"`
}

var TemplateModel, _ = modelImplementation.ModelFactory("mongo", config.Client, "template", "dbName")
