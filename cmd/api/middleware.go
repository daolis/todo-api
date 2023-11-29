package api

import (
	"net/http"
)

func (a *app) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// RFC 7807 (error returns)
		//defer func() {
		//	if err := recover(); err != nil {
		//		w.Header().Set("Connection", "close")
		//		a.logger.Error(err.(string))
		//		w.WriteHeader(http.StatusInternalServerError)
		//	}
		//}()

		next.ServeHTTP(w, r)
	})
}

func (a *app) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}
