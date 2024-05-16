package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares/logging"
	"httpServer/src/models"
	"net/http"
	"strconv"
)

type ApiInterface interface {
	Listen()
	Initialisation()
}

type Api struct {
	Json initialisation.JsonHandler
}

func (a Api) Listen() {
	var configuration models.Configuration
	if !a.Initialisation(&configuration) {
		return
	}
	r := mux.NewRouter()
	r.Use(logging.Logging())
	controller.InitControllers(r)
	fmt.Println("Server", configuration.Name, "starts listening on port:", configuration.Port)
	http.ListenAndServe(":"+strconv.Itoa(configuration.Port), r)
}

func (a Api) Initialisation(configuration *models.Configuration) bool {
	if !a.Json.ReadFile(configuration) {
		return false
	}
	fmt.Println("port:", configuration.Port)
	fmt.Println("name:", configuration.Name)
	fmt.Println("Database:")
	fmt.Println("\turl:", configuration.Db.Url)
	fmt.Println("\tname:", configuration.Db.Name)
	fmt.Println("\tport:", configuration.Db.Port)
	fmt.Println("\tuser:", configuration.Db.User)
	fmt.Println("\tpassword:", configuration.Db.Password)
	fmt.Println("data models:")
	fmt.Println("total:", len(configuration.Models))
	for _, model := range configuration.Models {
		fmt.Println("\tname:", model.Name)
		fmt.Print("\tfields:", len(model.Fields), " ")
		for _, e := range model.Fields {
			fmt.Print(e.Value + " ")
		}
		fmt.Println()
		fmt.Println("\tcreate:", model.Create)
		fmt.Println("\tread one:", model.ReadOne)
		fmt.Println("\tread many:", model.ReadMany)
		fmt.Println("\tupdate:", model.Update)
		fmt.Println("\tdelete:", model.Delete)
		fmt.Println("")
	}
	return true
}

type ApiService struct {
	Api *Api
}

func (s *ApiService) Listen() {
	s.Api.Listen()
}
