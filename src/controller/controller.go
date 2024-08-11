package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller/create"
	"httpServer/src/controller/deleteResource"
	"httpServer/src/controller/health"
	"httpServer/src/controller/patch"
	"httpServer/src/controller/readMany"
	"httpServer/src/controller/readOne"
	"httpServer/src/controller/search"
	"httpServer/src/controller/update"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/models"
	"net/http"
)

func InitControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel, db database2.DatabaseInterface) {
	health.InitController(r, dataModel)
	initCustomControllers(r, configuration, dataModel, db)
}

func test(w http.ResponseWriter, r *http.Request) {
	endpointLength := 5 // /test
	fmt.Println(r.RequestURI)
	fmt.Fprintf(w, "Custom controller "+r.RequestURI[:len(r.RequestURI)-endpointLength]+" is working\n")
}

func initCustomControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel, db database2.DatabaseInterface) {
	for _, field := range configuration.Models {
		controller := r.PathPrefix("/" + field.Name).Subrouter()
		var d initialisation.DataModel

		fmt.Println("Initializing", "/"+field.Name, "controller")
		controller.HandleFunc("/test", test).Methods("GET")
		endpointInitializers := map[*bool]func(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface){
			&field.Create:   create.InitCreateEndpoint,
			&field.ReadOne:  readOne.InitReadOneEndpoint,
			&field.ReadMany: readMany.InitReadManyEndpoint,
			&field.Update:   update.InitUpdateEndpoint,
			&field.Delete:   deleteResource.InitDeleteEndpoint,
			&field.Patch:    patch.InitPatchEndpoint,
			&field.Search:   search.InitSearchEndpoint,
		}
		for i, elem := range *dataModel {
			if elem.Name == field.Name {
				d = (*dataModel)[i]
			}
		}
		for condition, initializer := range endpointInitializers {
			if *condition {
				initializer(controller, d, db)
			}
		}
		fmt.Println()
	}
}
