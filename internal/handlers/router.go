package handlers

import (
	"es_load_test/appcontext"
	esLoad "es_load_test/internal/handlers/es_load"
	"net/http"

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
	router.Handle("/es_load/simple/item", getItem).Methods(http.MethodGet)

	postItem := esLoad.CreateItemHandler(s.Item)
	router.Handle("/es_load/simple/item", postItem).Methods(http.MethodPost)
}
