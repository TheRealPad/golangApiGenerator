package search

import (
	"fmt"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"net/http"
)

func InitSearchEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		lst, err := db.ReadMany(dataModel)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
		} else {
			if lst == nil {
				utils.HttpResponse(make([]int, 0), w, http.StatusOK)
			} else {
				lst = utils.Rsql(r.URL.Query().Get("rsql"), lst)
				if lst == nil {
					utils.HttpResponse(make([]int, 0), w, http.StatusOK)
				} else {
					utils.HttpResponse(lst, w, http.StatusOK)
				}
			}
		}
	}).Methods("GET")
	fmt.Println("init /search endpoint............................OK")
}
