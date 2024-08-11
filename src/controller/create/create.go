package create

import (
	"fmt"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitCreateEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !utils.GetRequestData(false, &d, w, r, true) {
			return
		}
		newData, err := db.Create(d)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			utils.HttpResponse(newData, w, http.StatusCreated)
		}
	}).Methods("POST")
	fmt.Println("init /create endpoint...........................OK")
}
