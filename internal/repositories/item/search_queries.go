package item

import (
	"context"
	"encoding/json"
	"es_load_test/internal/models"
	"es_load_test/internal/repositories/search"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

func (s *QueriesService) SearchQuery() SearchQueryFunc {
	return s.searchResult
}

func (s *QueriesService) esSearch(ctx context.Context, req search.Request) (*elastic.SearchResult, search.GetHits, error) {
	elasticQ, getHits := s.qs.BuildSearchQuery(req)

	searchSvc := s.client.Search().
		Index(req.IndexName).
		Query(elasticQ)

	var err error
	results, err := searchSvc.
		From(req.FromIndex()).Size(req.Size).Do(ctx)

	return results, getHits, err
}

func (s *QueriesService) searchResult(ctx context.Context, req search.Request) (search.Result, error) {
	results, getHits, err := s.esSearch(ctx, req)
	if err != nil {
		return search.Result{}, errors.Wrap(err, "without aggregations")
	}

	hits, err := getHits(results)
	if err != nil {
		return search.Result{}, errors.Wrap(err, "Unable to get hits for item")

	}

	var items []models.Item

	for _, hit := range hits {
		var i models.Item
		err = json.Unmarshal(hit.Source, &i)
		if err != nil {
			continue
		}
		i.Score = hit.Score
		items = append(items, i)
	}

	sr := search.Result{Items: items}

	return sr, nil
}
