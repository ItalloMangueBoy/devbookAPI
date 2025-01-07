package middlewares

import (
	"devbookAPI/src/auth"
	"devbookAPI/src/view"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidToken(r); err != nil {
			view.GenErrorTemplate(err).Send(w, 401)
			return
		}

		next(w, r)
	}
}
