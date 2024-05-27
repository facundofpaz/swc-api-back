package webresponse

import (
	"encoding/json"
	"net/http"
)

func NewWebResponse[T any](body T, w http.ResponseWriter, statusCode int) {
	data, _ := json.Marshal(body)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
