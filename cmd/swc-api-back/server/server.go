package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()

	urlMappings := newMapping()
	urlMappings.mapUrlsToHandlers(r)

	http.ListenAndServe(":8080", r)
}
