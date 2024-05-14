package logging

import (
	"httpServer/src/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Logs []models.Log

func Logging() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Logs = append(Logs, models.Log{Method: r.Method, Url: r.URL.Path, Address: r.RemoteAddr, Time: time.Now()})
			defer func() { log.Println(r.Method, r.URL.Path, r.RemoteAddr) }()
			next.ServeHTTP(w, r)
		})
	}
}
