package search

import (
	"es_load_test/internal/models"
)

type Result struct {
	Items        []models.Item
	Aggregations []string
}
