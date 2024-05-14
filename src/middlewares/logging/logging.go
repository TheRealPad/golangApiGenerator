package logging

import (
	"httpServer/src/middlewares"
	"httpServer/src/models"
	"log"
	"net/http"
	"time"
)

var Logs []models.Log

func Logging() middlewares.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			Logs = append(Logs, models.Log{Method: r.Method, Url: r.URL.Path, Address: r.RemoteAddr, Time: time.Now()})
			defer func() { log.Println(r.Method, r.URL.Path, r.RemoteAddr) }()
			f(w, r)
		}
	}
}
