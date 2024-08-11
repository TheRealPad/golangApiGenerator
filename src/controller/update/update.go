package update

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitUpdateEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/update/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		if !utils.GetRequestData(false, &d, w, r, true) {
			return
		}
		vars := mux.Vars(r)
		pathUuid := vars["uuid"]
		parseUuid, _ := uuid.Parse(pathUuid)
		d.Fields["uuid"].SetData(parseUuid.String(), initialisation.Uuid)
		_, err := db.Update(parseUuid, d)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			utils.HttpResponse(d.Fields, w, http.StatusOK)
		}
	}).Methods("PUT")
	fmt.Println("init /update endpoint...........................OK")
}
