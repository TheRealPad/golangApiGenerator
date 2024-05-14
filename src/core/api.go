package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"net/http"
	"strconv"
)

type ApiInterface interface {
	Listen()
}

type Api struct {
	Port int
}

func (a Api) Listen() {
	r := mux.NewRouter()
	controller.InitControllers(r)
	fmt.Print("Start listening on port: " + strconv.Itoa(a.Port) + "\n")
	http.ListenAndServe(":"+strconv.Itoa(a.Port), r)
}

type ApiService struct {
	Api *Api
}

func (s *ApiService) Listen() {
	s.Api.Listen()
}
