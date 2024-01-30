package publication

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"swc-api-back.com/internal/domain/model/publication"
	publicationSrv "swc-api-back.com/internal/service/publication"
)

type Handler struct {
	publicationService publicationSrv.Publication
}

func New(publicationService publicationSrv.Publication) (*Handler, error) {
	if publicationService == nil {
		return nil, errors.New("missing publication service")
	}
	return &Handler{
		publicationService: publicationService,
	}, nil
}

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idU, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	p, err := h.publicationService.GetByID(r.Context(), publication.ID(idU))
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, _ := json.Marshal(p)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
