package utils

import (
	"encoding/json"
	"fmt"
	"httpServer/src/initialisation"
	"io"
	"net/http"
)

func getKey(d *initialisation.DataModel, key string, requestData interface{}, w http.ResponseWriter) bool {
	value, ok := requestData.(map[string]interface{})[key]
	if !ok {
		fmt.Printf("Key %s not found in JSON data\n", key)
		HttpResponse(map[string]string{"error": "missing field in request body: " + key}, w, http.StatusBadRequest)
		return false
	}
	d.Fields[key].SetData(value.(string), d.Fields[key].GetDataType())
	return true
}

func GetRequestData(getUuid bool, d *initialisation.DataModel, w http.ResponseWriter, r *http.Request, allFields bool) bool {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		HttpResponse(map[string]string{"error": "Failed to read request body"}, w, http.StatusBadRequest)
		return false
	}
	var requestData interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		HttpResponse(map[string]string{"error": "Failed to parse JSON body"}, w, http.StatusBadRequest)
		return false
	}
	if !getUuid {
		d.Fields[initialisation.Uuid].SetData(GenerateUuid(), initialisation.Uuid)
	}
	if !allFields {
		return true
	}
	for key := range d.Fields {
		if (key != initialisation.Uuid || getUuid && key == initialisation.Uuid) && !getKey(d, key, requestData, w) {
			return false
		}
	}
	return true
}
