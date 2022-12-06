package item

import (
	"context"
	"es_load_test/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"es_load_test/config"
	"es_load_test/internal/repositories/cluster"
	"es_load_test/internal/repositories/index"
	"es_load_test/internal/repositories/search"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

type SearchQueryFunc func(ctx context.Context, req search.Request) (search.Result, error)
type CreateIndexQueryFunc func(ctx context.Context, req cluster.Request) error
type IndexDocQueryFunc func(ctx context.Context, req index.Request) (models.Item, error)
type BulkIndexDocQueryFunc func(ctx context.Context, req index.Request) error

type Queries interface {
	SearchQuery() SearchQueryFunc
	CreateIndexQuery() CreateIndexQueryFunc
	IndexDocQuery() IndexDocQueryFunc
	BulkIndexDocQuery() BulkIndexDocQueryFunc
}

type QueriesService struct {
	client *elastic.Client
	config config.Elasticsearch
	qs     BuilderService
}

func NewQueriesService(queryService BuilderService) (*QueriesService, error) {
	esCfg := config.Cfg.GetES()
	c := http.Client{
		Timeout: time.Duration(esCfg.HTTPTimeoutMS) * time.Second,
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHttpClient(&c),
		elastic.SetURL(esCfg.Hosts...),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetTraceLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)), // uncomment this if you want to print query
	)

	if err != nil {
		return nil, errors.Wrap(err, "couldn't create elastic search client for item search")
	}

	return &QueriesService{
		client: client,
		config: esCfg,
		qs:     queryService,
	}, nil
}
