package middleware

import (
	"chaobum-api/config"
	"net/http"
)

func SetCors(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.Env.FRONT_URL)
		//pre-flight処理を可能にするために"Content-Type"ヘッダーを許可
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		//pre-flight処理を可能にするためにOPTIONSメソッドの場合はreturn
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	}
}
