package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares/logging"
	"net/http"
	"strconv"
)

type ApiInterface interface {
	Listen()
	Initialisation()
}

type Api struct {
	Port int
	Json initialisation.JsonHandler
}

func (a Api) Listen() {
	a.Initialisation()
	r := mux.NewRouter()
	r.Use(logging.Logging())
	controller.InitControllers(r)
	fmt.Print("Start listening on port: " + strconv.Itoa(a.Port) + "\n")
	http.ListenAndServe(":"+strconv.Itoa(a.Port), r)
}

func (a Api) Initialisation() {
	a.Json.ReadFile()
}

type ApiService struct {
	Api *Api
}

func (s *ApiService) Listen() {
	s.Api.Listen()
}
