package api

import (
	"net/http"
)

func (a *app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := a.writeJSON(w, status, message)
	if err != nil {
		a.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *app) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	a.errorResponse(w, r, http.StatusNotFound, message)
}
