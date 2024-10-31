package view

import "net/http"

type errorTemplate struct {
	Msg string `json:"error"`
}

// GenErrorTemplate: generete an generic error response
func GenErrorTemplate(err error) errorTemplate {
	return errorTemplate{Msg: err.Error()}
}

// Send: sends JSON error response
func (err errorTemplate) Send(w http.ResponseWriter, statusCode int) {
	JSON(w, statusCode, err)
}
