package errorweb

import (
	"encoding/json"
	"net/http"
)

type (
	webError struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}
	NotFoundErr struct {
		Message string
	}
	BadRequestErr struct {
		Message string
	}
)

func (e NotFoundErr) Error() string {
	return e.Message
}

func (e BadRequestErr) Error() string {
	return e.Message
}

func NewWebError(e error, w http.ResponseWriter) {
	var err webError
	switch e.(type) {
	case NotFoundErr:
		err.StatusCode = http.StatusNotFound
	case BadRequestErr:
		err.StatusCode = http.StatusBadRequest
	default:
		err.StatusCode = http.StatusInternalServerError
	}
	err.Message = e.Error()
	dataError, _ := json.Marshal(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	w.Write(dataError)
}
