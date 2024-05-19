package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"httpServer/src/controller/health"
	"httpServer/src/initialisation"
	"httpServer/src/models"
	"httpServer/src/utils"
	"io"
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

func jsonResponse(data interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal response to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
	fmt.Printf("%d - ", statusCode)
}

func getKey(d *initialisation.DataModel, key string, requestData interface{}, w http.ResponseWriter) bool {
	value, ok := requestData.(map[string]interface{})[key]
	if !ok {
		fmt.Printf("Key %s not found in JSON data\n", key)
		jsonResponse(map[string]string{"error": "missing field in request body: " + key}, w, http.StatusBadRequest)
		return false
	}
	d.Fields[key].SetData(value.(string), d.Fields[key].GetDataType())
	return true
}

func getRequestData(getUuid bool, d *initialisation.DataModel, w http.ResponseWriter, r *http.Request) bool {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return false
	}
	var requestData interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return false
	}
	if !getUuid {
		d.Fields[initialisation.Uuid].SetData(utils.GenerateUuid(), initialisation.Uuid)
	}
	for key := range d.Fields {
		if (key != initialisation.Uuid || getUuid && key == initialisation.Uuid) && !getKey(d, key, requestData, w) {
			return false
		}
	}
	return true
}

func initCreateEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !getRequestData(false, &d, w, r) {
			return
		}
		jsonResponse(d.Fields, w, http.StatusCreated)
	}).Methods("POST")
	fmt.Println("init /create endpoint...........................OK")
}

func initReadOneEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	r.HandleFunc("/readOne", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		jsonResponse(d.Fields, w, http.StatusOK)
	}).Methods("GET")
	fmt.Println("init /read one endpoint.........................OK")
}

func initReadManyEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		var lst []initialisation.Field
		lst = append(lst, d.Fields)
		lst = append(lst, d.Fields)
		lst = append(lst, d.Fields)
		jsonResponse(lst, w, http.StatusOK)
	}).Methods("GET")
	fmt.Println("init /read many endpoint........................OK")
}

func initUpdateEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	r.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !getRequestData(false, &d, w, r) {
			return
		}
		jsonResponse(d.Fields, w, http.StatusOK)
	}).Methods("PUT")
	fmt.Println("init /update endpoint...........................OK")
}

func initDeleteEndpoint(r *mux.Router, dataModel initialisation.DataModel) {
	r.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		jsonResponse(d.Fields, w, http.StatusNoContent)
	}).Methods("DELETE")
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
		var d initialisation.DataModel
		for i, elem := range *dataModel {
			if elem.Name == field.Name {
				d = (*dataModel)[i]
			}
		}
		for condition, initializer := range endpointInitializers {
			if *condition {
				initializer(controller, d)
			}
		}
		fmt.Println()
	}
}
