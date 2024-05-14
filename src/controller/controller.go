package controller

import (
	"github.com/gorilla/mux"
	"httpServer/src/controller/health"
)

func InitControllers(r *mux.Router) {
	health.InitController(r)
}
