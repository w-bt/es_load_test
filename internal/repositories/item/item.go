package item

import (
	"context"

	"es_load_test/internal/models"
	"es_load_test/internal/repositories/cluster"
	"es_load_test/internal/repositories/index"
	"es_load_test/internal/repositories/search"
)

type Repository interface {
	Search(ctx context.Context, req search.Request) (search.Result, error)
	CreateIndex(ctx context.Context, req cluster.Request) error
	IndexDoc(ctx context.Context, req index.Request) (models.Item, error)
	BulkIndexDoc(ctx context.Context, req index.Request) error
}

type itemRepo struct {
	queries Queries
}

func NewItemRepo(queries Queries) Repository {
	return &itemRepo{queries: queries}
}

func (repo itemRepo) Search(ctx context.Context, req search.Request) (search.Result, error) {
	query := repo.queries.SearchQuery()
	result, err := query(ctx, req)
	if err != nil {
		return search.Result{}, err
	}

	return result, nil
}

func (repo itemRepo) CreateIndex(ctx context.Context, req cluster.Request) error {
	query := repo.queries.CreateIndexQuery()
	err := query(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (repo itemRepo) IndexDoc(ctx context.Context, req index.Request) (models.Item, error) {
	query := repo.queries.IndexDocQuery()
	item, err := query(ctx, req)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (repo itemRepo) BulkIndexDoc(ctx context.Context, req index.Request) error {
	query := repo.queries.BulkIndexDocQuery()
	err := query(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
