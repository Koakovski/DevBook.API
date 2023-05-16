package middlwares

import (
	presenter "devbook-api/src/app/presenters"
	"devbook-api/src/infra/auth"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			presenter.ErrorPresenter(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
