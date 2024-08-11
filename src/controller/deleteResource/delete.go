package deleteResource

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitDeleteEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/deleteResource/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		vars := mux.Vars(r)
		pathUuid := vars["uuid"]
		parseUuid, _ := uuid.Parse(pathUuid)
		_, err := db.Delete(parseUuid, dataModel.Name)
		if err != nil {
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			utils.HttpResponse(d.Fields, w, http.StatusNoContent)
		}
	}).Methods("DELETE")
	fmt.Println("init /deleteResource endpoint...........................OK")
}
