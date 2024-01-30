package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication"
)

type mapping struct {
	publicationHandler *publication.Handler
}

func newMapping() mapping {
	return mapping{
		publicationHandler: resolvePublicationHandler(),
	}
}

func (m mapping) mapUrlsToHandlers(r *mux.Router) {
	r.HandleFunc("/swc-api/publication/{id}", m.publicationHandler.GetByID).Methods(http.MethodGet)
}
