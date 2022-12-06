package item

import (
	"context"
	"es_load_test/internal/repositories/cluster"
	"fmt"
)

func (s *QueriesService) CreateIndexQuery() CreateIndexQueryFunc {
	return s.createIndexResult
}

func (s *QueriesService) createIndexResult(ctx context.Context, req cluster.Request) error {
	exist, err := s.client.IndexExists(req.IndexName).Do(ctx)
	if err != nil {
		return err
	}

	if exist {
		fmt.Println("The index " + req.IndexName + " already exists.")
		_, err = s.client.DeleteIndex(req.IndexName).Do(ctx)
		if err != nil {
			return err
		}
	}

	create, err := s.client.CreateIndex(req.IndexName).Body(mappings).Do(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Create index:", create)

	return nil
}
