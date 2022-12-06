package item

import (
	"context"

	"es_load_test/internal/repositories/index"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

func (s *QueriesService) BulkIndexDocQuery() BulkIndexDocQueryFunc {
	return s.bulkIndexDocResult
}

func (s *QueriesService) bulkIndexDocResult(ctx context.Context, req index.Request) error {
	exist, err := s.client.IndexExists(req.IndexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("index not found")
	}

	bulk := s.client.Bulk()

	for i := 0; i < req.BatchSize; i++ {
		item := s.qs.GenerateDoc()

		r := elastic.NewBulkIndexRequest()
		r.OpType("index") // set type to "index" document
		r.Index(req.IndexName)
		r.Id(item.Code)
		r.Doc(item)

		bulk = bulk.Add(r)
	}

	_, err = bulk.Do(ctx)
	if err != nil {
		return err
	}

	if req.Refresh {
		_, err = s.client.Refresh(req.IndexName).Do(ctx)
		if err != nil {
			return err
		}
	}

	if req.Flush {
		_, err = s.client.Flush(req.IndexName).Do(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
