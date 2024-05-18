package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller"
	"httpServer/src/initialisation"
	"httpServer/src/middlewares"
	"httpServer/src/models"
	"net/http"
	"strconv"
	"strings"
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
	if !a.Initialisation(&configuration, &dataModel) {
		return
	}
	displayDataTypes(&dataModel)
	r := mux.NewRouter()
	middlewares.GlobalMiddleware(r)
	controller.InitControllers(r)
	fmt.Println("Server", configuration.Name, "starts listening on port:", configuration.Port)
	http.ListenAndServe(":"+strconv.Itoa(configuration.Port), r)
}

func (a Api) Initialisation(configuration *models.Configuration, dataModel *[]initialisation.DataModel) bool {
	if !a.Json.ReadFile(configuration) {
		return false
	}
	for _, model := range configuration.Models {
		*dataModel = append(*dataModel, initialisation.DataModel{Name: model.Name, Fields: make(initialisation.Field)})
		dataModelPtr := &(*dataModel)[len(*dataModel)-1]
		for _, e := range model.Fields {
			separator := " - "
			parts := strings.SplitN(e.Value, separator, 2)
			dataModelPtr.Fields[parts[0]] = &initialisation.DynamicType{}
			dataModelPtr.Fields[parts[0]].SetData("", initialisation.Datatype(parts[1]))
		}
	}
	displayConfiguration(configuration)
	return true
}

func displayConfiguration(configuration *models.Configuration) {
	fmt.Println("CONFIGURATION:\n")
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
}

func displayDataTypes(dataModel *[]initialisation.DataModel) {
	fmt.Println("DATA TYPES:\n")
	for _, elem := range *dataModel {
		fmt.Println(elem.Name)
		for k, f := range elem.Fields {
			fmt.Println("\t", k, ":", f.GetDataType())
		}
	}
}

type ApiService struct {
	Api *Api
}

func (s *ApiService) Listen() {
	s.Api.Listen()
}
