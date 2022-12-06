package item

import (
	"es_load_test/internal/models"
	"es_load_test/internal/repositories/search"

	"github.com/olivere/elastic/v7"
)

type BuilderService interface {
	BuildSearchQuery(req search.Request) (elastic.Query, search.GetHits)
	GenerateDoc() *models.Item
}

type builderService struct{}

func NewBuilderService() BuilderService {
	return &builderService{}
}
