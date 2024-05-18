package middlewares

import (
	"github.com/gorilla/mux"
	"httpServer/src/middlewares/logging"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func GlobalMiddleware(r *mux.Router) {
	r.Use(logging.Logging())
}
