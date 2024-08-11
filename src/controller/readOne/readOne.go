package readOne

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitReadOneEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/read/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["uuid"]
		ok, _ := uuid.Parse(id)
		fields, err := db.ReadOne(ok, dataModel)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else if fields == nil {
			utils.HttpResponse(map[string]string{"message": "no data"}, w, http.StatusNotFound)
		} else {
			utils.HttpResponse(fields, w, http.StatusOK)
		}
	}).Methods("GET")
	fmt.Println("init /read one endpoint.........................OK")
}
