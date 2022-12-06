package es_load

import (
	"net/url"
	"strconv"
)

const (
	defaultSize = 20
)

func getQuery(query url.Values) string {
	return query.Get("query")
}

func getIndex(query url.Values) string {
	return query.Get("index")
}

func getRefresh(query url.Values) bool {
	return query.Get("refresh") == "true"
}

func getPageNumber(query url.Values) int {
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		return 0
	}

	return page
}

func getPageSize(query url.Values) int {
	size, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		return defaultSize
	}

	return size
}
