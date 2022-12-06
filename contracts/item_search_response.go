package contracts

import "es_load_test/internal/models"

type SearchResponse struct {
	Items        []models.Item `json:"items"`
	Aggregations []string      `json:"aggregations,omitempty"`
}
