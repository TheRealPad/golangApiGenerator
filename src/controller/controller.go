package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"httpServer/src/controller/health"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/models"
	"httpServer/src/utils"
	"io"
	"net/http"
)

func InitControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel, db database2.DatabaseInterface) {
	health.InitController(r)
	initCustomControllers(r, configuration, dataModel, db)
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
		jsonResponse(map[string]string{"error": "Failed to read request body"}, w, http.StatusBadRequest)
		return false
	}
	var requestData interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		jsonResponse(map[string]string{"error": "Failed to parse JSON body"}, w, http.StatusBadRequest)
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

func initCreateEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !getRequestData(false, &d, w, r) {
			return
		}
		newData, err := db.Create(d)
		if err != nil {
			jsonResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			jsonResponse(newData, w, http.StatusCreated)
		}
	}).Methods("POST")
	fmt.Println("init /create endpoint...........................OK")
}

func initReadOneEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/read/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["uuid"]
		ok, _ := uuid.Parse(id)
		fields, err := db.ReadOne(ok, dataModel)
		if err != nil {
			jsonResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else if fields == nil {
			jsonResponse(map[string]string{"message": "no data"}, w, http.StatusNotFound)
		} else {
			jsonResponse(fields, w, http.StatusOK)
		}
	}).Methods("GET")
	fmt.Println("init /read one endpoint.........................OK")
}

func initReadManyEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		lst, err := db.ReadMany(dataModel)
		if err != nil {
			jsonResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			jsonResponse(lst, w, http.StatusOK)
		}
	}).Methods("GET")
	fmt.Println("init /read many endpoint........................OK")
}

func initUpdateEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/update/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !getRequestData(false, &d, w, r) {
			return
		}
		vars := mux.Vars(r)
		uuid := vars["uuid"]
		d.Fields[initialisation.Uuid].SetData(uuid, initialisation.Uuid)
		jsonResponse(d.Fields, w, http.StatusOK)
	}).Methods("PUT")
	fmt.Println("init /update endpoint...........................OK")
}

func initDeleteEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/delete/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		vars := mux.Vars(r)
		pathUuid := vars["uuid"]
		parseUuid, _ := uuid.Parse(pathUuid)
		_, err := db.Delete(parseUuid, dataModel.Name)
		if err != nil {
			jsonResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			jsonResponse(d.Fields, w, http.StatusNoContent)
		}
	}).Methods("DELETE")
	fmt.Println("init /delete endpoint...........................OK")
}

func initCustomControllers(r *mux.Router, configuration *models.Configuration, dataModel *[]initialisation.DataModel, db database2.DatabaseInterface) {
	for _, field := range configuration.Models {
		controller := r.PathPrefix("/" + field.Name).Subrouter()
		fmt.Println("Initializing", "/"+field.Name, "controller")
		controller.HandleFunc("/test", test).Methods("GET")
		endpointInitializers := map[*bool]func(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface){
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
				initializer(controller, d, db)
			}
		}
		fmt.Println()
	}
}
