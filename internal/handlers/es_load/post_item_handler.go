package es_load

import (
	"log"
	"net/http"

	"es_load_test/constant"
	"es_load_test/internal/repositories/index"
	"es_load_test/internal/services/item"
	"es_load_test/internal/utils"
)

func CreateItemHandler(is item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		esReq := createIndexRequest(req)
		userData, err := is.IndexDoc(ctx, esReq)
		if err != nil {
			log.Printf("Error while creating item : %s", err)
			utils.WriteResponse(w, http.StatusInternalServerError, nil)
			return
		}

		log.Printf("Successfully create random item\n")
		utils.WriteResponse(w, http.StatusOK, userData)
		return
	}
}

var createIndexRequest = func(r *http.Request) index.Request {
	qParam := r.URL.Query()
	indexName := getIndex(qParam)
	if indexName == "" {
		indexName = constant.DefaultIndexName
	}
	refresh := getRefresh(qParam)

	return index.Request{
		IndexName: indexName,
		Refresh:   refresh,
	}
}
