package middleware

import (
	"chaobum-api/config"
	"net/http"
)

func SetCors(handler http.Handler) http.Handler {
	corsSetter := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.Env.FRONT_URL)

		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(corsSetter)
}
