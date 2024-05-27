package publication

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	errorweb "swc-api-back.com/cmd/swc-api-back/handlers/error"
	webresponse "swc-api-back.com/cmd/swc-api-back/handlers/web-response"
	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication/dto"
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

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var publicationRequest dto.Publication
	err := json.NewDecoder(r.Body).Decode(&publicationRequest)
	if err != nil {
		errorweb.NewWebError(errorweb.BadRequestErr{Message: "error on decode creation publication body"}, w)
		return
	}
	publicationID, err := h.publicationService.Create(r.Context(), dto.ToDomain(publicationRequest))
	if err != nil {
		errorweb.NewWebError(err, w)
		return
	}
	webresponse.NewWebResponse(dto.Publication{ID: publicationID}, w, http.StatusCreated)
}
