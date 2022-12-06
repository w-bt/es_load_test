package contracts

type BulkIndexDocResponse struct {
	Message  string `json:"message"`
	TotalDoc int    `json:"total_doc"`
}
