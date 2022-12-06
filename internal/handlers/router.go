package handlers

import (
	"net/http"

	"es_load_test/appcontext"
	esLoad "es_load_test/internal/handlers/es_load"

	"github.com/gorilla/mux"
)

func NewRouter(deps *appcontext.ServerDependencies) http.Handler {
	router := mux.NewRouter()

	router.UseEncodedPath()
	registerAPI(deps, router)

	newRouter := withDefaultResponseHeaders(router)
	return http.HandlerFunc(newRouter.ServeHTTP)
}

func registerAPI(deps *appcontext.ServerDependencies, router *mux.Router) {
	s := deps.Services

	getItem := esLoad.GetItemHandler(s.Item)
	router.Handle("/es_load/item", getItem).Methods(http.MethodGet)

	postItem := esLoad.CreateItemHandler(s.Item)
	router.Handle("/es_load/item", postItem).Methods(http.MethodPost)

	postBulkItem := esLoad.CreateBulkItemHandler(s.Item)
	router.Handle("/es_load/bulk_item", postBulkItem).Methods(http.MethodPost)
}
