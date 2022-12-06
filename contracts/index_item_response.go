package contracts

import "es_load_test/internal/models"

type IndexDocResponse struct {
	Message string      `json:"message"`
	Data    models.Item `json:"data"`
}
