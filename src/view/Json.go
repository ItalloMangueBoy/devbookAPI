package view

import (
	"encoding/json"
	"net/http"
)

// JSON: sends JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(500)
		GenErrorTemplate(err).Send(w, statusCode)
	}

}
