package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"httpServer/src/controller/health"
	"httpServer/src/initialisation"
	"httpServer/src/models"
	"io/ioutil"
	"net/http"
)

func InitControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel) {
	health.InitController(r)
	initCustomControllers(r, configuration, dataModel)
}

func test(w http.ResponseWriter, r *http.Request) {
	size := 5
	fmt.Fprintf(w, "Custom controller "+r.RequestURI[:len(r.RequestURI)-size]+" is working\n")
}

func initCreateEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	fmt.Println("init /create endpoint...........................OK")
	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		var requestData interface{}
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}
		d := dataModel
		fmt.Fprintf(w, "{\n")
		d.Fields["uuid"].SetData(uuid.New().String(), initialisation.Uuid)
		fmt.Fprintf(w, "\t\"uuid\": %+v\n", d.Fields["uuid"].GetData())
		for key, _ := range d.Fields {
			if key != "uuid" {
				value, ok := requestData.(map[string]interface{})[key]
				if !ok {
					fmt.Printf("Key %s not found in JSON data\n", key)
					continue
				}
				d.Fields[key].SetData(value.(string), d.Fields[key].GetDataType())
				fmt.Printf("%s : %s - %s\n", key, value, d.Fields[key].GetDataType())
				fmt.Fprintf(w, "\t\"%s\": %+v\n", key, d.Fields[key].GetData())
			}
		}
		fmt.Fprintf(w, "}\n")
	}).Methods("POST")
}

func initReadOneEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	fmt.Println("init /read one endpoint.........................OK")
}

func initReadManyEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	fmt.Println("init /read many endpoint........................OK")
}

func initUpdateEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	fmt.Println("init /update endpoint...........................OK")
}

func initDeleteEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	fmt.Println("init /delete endpoint...........................OK")
}

func initCustomControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel) {
	for _, field := range configuration.Models {
		controller := r.PathPrefix("/" + field.Name).Subrouter()
		fmt.Println("Initializing", "/"+field.Name, "controller")
		controller.HandleFunc("/test", test).Methods("GET")
		endpointInitializers := map[*bool]func(r *mux.Router, dataModel initialisation.DataModel){
			&field.Create:   initCreateEndpoint,
			&field.ReadOne:  initReadOneEndpoint,
			&field.ReadMany: initReadManyEndpoint,
			&field.Update:   initUpdateEndpoint,
			&field.Delete:   initDeleteEndpoint,
		}
		for condition, initializer := range endpointInitializers {
			if *condition {
				for i, elem := range *dataModel {
					if elem.Name == field.Name {
						initializer(controller, (*dataModel)[i])
					}
				}
			}
		}
		fmt.Println()
	}
}
