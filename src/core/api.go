package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares"
	"httpServer/src/middlewares/apiKey"
	"httpServer/src/models"
	"httpServer/src/utils"
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
	var dataModel []initialisation.DataModel
	var db database2.DatabaseInterface
	if !a.Initialisation(&configuration, &dataModel) {
		return
	}
	if configuration.IsSecure {
		apiKey.ApiKey, _ = utils.GenerateKey(100)
		fmt.Println("Api key: ", apiKey.ApiKey)
	}
	db = &database2.MongoDB{Name: configuration.Db.Name, Url: configuration.Db.Url}
	displayDataTypes(&dataModel)
	r := mux.NewRouter()
	middlewares.GlobalMiddleware(r, configuration.IsSecure)
	controller.InitControllers(r, &configuration, &dataModel, db)
	fmt.Println("Server", configuration.Name, "starts listening on port:", configuration.Port)
	http.ListenAndServe(":"+strconv.Itoa(configuration.Port), r)
}

type ApiService struct {
	Api *Api
}

func (s *ApiService) Listen() {
	s.Api.Listen()
}
