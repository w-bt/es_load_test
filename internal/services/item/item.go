package item

import (
	"context"
	"errors"
	"fmt"

	"es_load_test/contracts"
	"es_load_test/internal/repositories/cluster"
	"es_load_test/internal/repositories/index"
	"es_load_test/internal/repositories/item"
	"es_load_test/internal/repositories/search"
)

const ErrItemNotFound = "Item not found"

type Service interface {
	GetItem(ctx context.Context, req search.Request) (*contracts.SearchResponse, error)
	CreateIndex(ctx context.Context, req cluster.Request) (*contracts.CreateIndexResponse, error)
	IndexDoc(ctx context.Context, req index.Request) (*contracts.IndexDocResponse, error)
	BulkIndexDoc(ctx context.Context, req index.Request) (*contracts.BulkIndexDocResponse, error)
}

type itemService struct {
	itemRepository item.Repository
}

func NewItemService(repo item.Repository) Service {
	return &itemService{
		itemRepository: repo,
	}
}

func (service itemService) GetItem(ctx context.Context, req search.Request) (*contracts.SearchResponse, error) {
	result, err := service.itemRepository.Search(ctx, req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while fetching item with req %+v", req))
	}

	if len(result.Items) == 0 {
		return nil, errors.New(ErrItemNotFound)
	}

	return &contracts.SearchResponse{
		Items: result.Items,
	}, nil
}

func (service itemService) CreateIndex(ctx context.Context, req cluster.Request) (*contracts.CreateIndexResponse, error) {
	err := service.itemRepository.CreateIndex(ctx, req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while fetching item with req %+v", req))
	}

	return &contracts.CreateIndexResponse{
		Message: "success",
	}, nil
}

func (service itemService) IndexDoc(ctx context.Context, req index.Request) (*contracts.IndexDocResponse, error) {
	itemDoc, err := service.itemRepository.IndexDoc(ctx, req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while indexing item with req %+v, err %+v", req, err))
	}

	return &contracts.IndexDocResponse{
		Message: "success",
		Data:    itemDoc,
	}, nil
}

func (service itemService) BulkIndexDoc(ctx context.Context, req index.Request) (*contracts.BulkIndexDocResponse, error) {
	err := service.itemRepository.BulkIndexDoc(ctx, req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while fetching item with req %+v, err %+v", req, err))
	}

	return &contracts.BulkIndexDocResponse{
		Message:  "success",
		TotalDoc: req.BatchSize,
	}, nil
}
