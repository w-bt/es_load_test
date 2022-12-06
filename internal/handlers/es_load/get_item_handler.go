package es_load

import (
	"log"
	"net/http"

	"es_load_test/constant"
	"es_load_test/internal/repositories/search"
	"es_load_test/internal/services/item"
	"es_load_test/internal/utils"

	"github.com/bxcodec/faker/v4"
)

func GetItemHandler(is item.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		esReq := createSearchRequest(req)
		r, err := is.GetItem(ctx, esReq)
		if err != nil {
			log.Printf("Error while getting latest item : %s", err)
			if err.Error() == item.ErrItemNotFound {
				utils.WriteResponse(w, http.StatusNotFound, nil)
				return
			}
			utils.WriteResponse(w, http.StatusInternalServerError, nil)
			return
		}

		log.Println("Successfully get latest item")
		utils.WriteResponse(w, http.StatusOK, r)
		return
	}
}

var createSearchRequest = func(r *http.Request) search.Request {
	qParam := r.URL.Query()
	query := getQuery(qParam)
	if query == "" {
		query = faker.Word()
	}
	page := getPageNumber(qParam)
	size := getPageSize(qParam)
	index := getIndex(qParam)
	if index == "" {
		index = constant.DefaultIndexName
	}
	agg := getAggregator(qParam)

	return search.Request{
		Query:     query,
		Agg:       agg,
		IndexName: index,
		Page:      page,
		NextPage:  page + 1,
		Size:      size,
	}
}
