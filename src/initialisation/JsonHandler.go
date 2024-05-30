package initialisation

import (
	"encoding/json"
	"httpServer/src/models"
	"io"
	"log"
	"os"
)

type JsonHandlerInterface interface {
	ReadFile()
}

type JsonHandler struct {
	File string
}

func (j JsonHandler) ReadFile(configuration *models.Configuration) bool {
	jsonFile, err := os.Open(j.File)
	if err != nil {
		log.Fatal(err)
		return false
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configuration)
	defer jsonFile.Close()
	return true
}
