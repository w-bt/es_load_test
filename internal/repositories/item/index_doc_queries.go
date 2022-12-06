package item

import (
	"context"
	"es_load_test/internal/models"
	"es_load_test/internal/repositories/index"
	"github.com/pkg/errors"
)

func (s *QueriesService) IndexDocQuery() IndexDocQueryFunc {
	return s.indexDocResult
}

func (s *QueriesService) indexDocResult(ctx context.Context, req index.Request) (models.Item, error) {
	exist, err := s.client.IndexExists(req.IndexName).Do(ctx)
	if err != nil {
		return models.Item{}, err
	}

	if !exist {
		return models.Item{}, errors.New("index not found")
	}

	item := s.qs.GenerateDoc()

	_, err = s.client.Index().
		Index(req.IndexName).
		BodyJson(item).
		Id(item.Code).
		Do(ctx)

	if req.Refresh {
		_, err = s.client.Refresh(req.IndexName).Do(ctx)
		if err != nil {
			return models.Item{}, err
		}
	}

	return *item, nil
}
