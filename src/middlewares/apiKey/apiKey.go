package apiKey

import (
	"github.com/gorilla/mux"
	"httpServer/src/utils"
	"net/http"
)

var ApiKey string

func CheckApiKey() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentApiKey := r.Header.Get("ApiKey")
			if len(currentApiKey) == 0 {
				utils.HttpResponse(map[string]string{"error": "ApiKey is missing in header"}, w, http.StatusUnauthorized)
				return
			}
			if currentApiKey != ApiKey {
				utils.HttpResponse(map[string]string{"error": "Bad value for ApiKey in header"}, w, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
