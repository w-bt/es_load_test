package search

import (
	"github.com/olivere/elastic/v7"
)

type GetHits func(results *elastic.SearchResult) (hits []*elastic.SearchHit, err error)

func DefaultGetHitsFunc(results *elastic.SearchResult) ([]*elastic.SearchHit, error) {
	return results.Hits.Hits, nil
}
