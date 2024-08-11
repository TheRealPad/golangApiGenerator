package patch

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	database2 "httpServer/src/database"
	"httpServer/src/initialisation"
	"httpServer/src/utils"
	"io"
	"net/http"
)

func InitPatchEndpoint(r *mux.Router, dataModel initialisation.DataModel, db database2.DatabaseInterface) {
	r.HandleFunc("/patch/{uuid}", func(w http.ResponseWriter, r *http.Request) {
		d := dataModel
		body, err := io.ReadAll(r.Body)
		vars := mux.Vars(r)
		pathUuid := vars["uuid"]
		parseUuid, _ := uuid.Parse(pathUuid)
		d.Fields["uuid"].SetData(parseUuid.String(), initialisation.Uuid)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Failed to read request body"}, w, http.StatusBadRequest)
			return
		}
		var requestData map[string]interface{}
		if err := json.Unmarshal(body, &requestData); err != nil {
			utils.HttpResponse(map[string]string{"error": "Failed to parse JSON body"}, w, http.StatusBadRequest)
			return
		}
		for key, val := range requestData {
			d.Fields[key].SetData(val.(string), d.Fields[key].GetDataType())
		}
		fields, err := db.ReadOne(parseUuid, dataModel)
		if err != nil {
			fmt.Println(err.Error())
			utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
			return
		} else if fields == nil {
			utils.HttpResponse(map[string]string{"message": "no data"}, w, http.StatusNotFound)
			return
		} else {
			for key := range d.Fields {
				_, ok := requestData[key]
				if !ok {
					d.Fields[key].SetData(fields[key].GetData().(string), d.Fields[key].GetDataType())
				}
			}
			_, err := db.Update(parseUuid, d)
			if err != nil {
				fmt.Println(err.Error())
				utils.HttpResponse(map[string]string{"error": "Internal server error"}, w, http.StatusInternalServerError)
			} else {
				utils.HttpResponse(d.Fields, w, http.StatusOK)
			}
		}
	}).Methods("PATCH")
	fmt.Println("init /patch endpoint............................OK")
}
