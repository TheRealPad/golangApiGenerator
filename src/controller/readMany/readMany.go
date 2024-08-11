package readMany

import (
	"fmt"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitReadManyEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		lst, err := db.ReadMany(dataModel)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			if lst == nil {
				utils.HttpResponse(make([]int, 0), w, http.StatusOK)
			} else {
				utils.HttpResponse(lst, w, http.StatusOK)
			}
		}
	}).Methods("GET")
	fmt.Println("init /read many endpoint........................OK")
}
