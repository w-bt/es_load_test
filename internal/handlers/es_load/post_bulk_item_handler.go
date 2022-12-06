package es_load

import (
	"log"
	"net/http"

	"es_load_test/constant"
	"es_load_test/internal/repositories/index"
	"es_load_test/internal/services/item"
	"es_load_test/internal/utils"
)

func CreateBulkItemHandler(is item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		esReq := createBulkIndexRequest(req)
		userData, err := is.BulkIndexDoc(ctx, esReq)
		if err != nil {
			log.Printf("Error while creating bulk item : %s", err)
			utils.WriteResponse(w, http.StatusInternalServerError, nil)
			return
		}

		log.Printf("Successfully create bulk random item\n")
		utils.WriteResponse(w, http.StatusOK, userData)
		return
	}
}

var createBulkIndexRequest = func(r *http.Request) index.Request {
	qParam := r.URL.Query()
	indexName := getIndex(qParam)
	if indexName == "" {
		indexName = constant.DefaultIndexName
	}
	refresh := getRefresh(qParam)
	flush := getFlush(qParam)
	batchSize := getBatchSize(qParam)
	if batchSize == 0 {
		batchSize = constant.DefaultBatchSize
	}

	return index.Request{
		IndexName: indexName,
		Refresh:   refresh,
		Flush:     flush,
		BatchSize: batchSize,
	}
}
