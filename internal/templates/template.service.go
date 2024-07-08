package templates

import (
	// "github.com/sreerag-rajan/boilerplate-go/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// templateService is a class that contains methods for template operations
type TemplateService struct {
	templateCollection *mongo.Collection
}

func (s *TemplateService) init() {
	// s.templateCollection = config.Database.Collection("templates")
}

func (s *TemplateService) GetTemplates() {}

func (s *TemplateService) CreaTetemplates(body interface{}) {}

func (s *TemplateService) GetTemplateById(id int) {}

func (s *TemplateService) DeleteTemplate(id int) {}

func (s *TemplateService) UpdateTemplate(id int, body interface{}) {}
