package core

import (
	"httpServer/src/initialisation"
	"httpServer/src/models"
)

func (a Api) initialisation(configuration *models.Configuration, dataModel *[]initialisation.DataModel) bool {
	if !a.Json.ReadFile(configuration) {
		return false
	}
	for _, model := range configuration.Models {
		*dataModel = append(*dataModel, initialisation.DataModel{Name: model.Name, Fields: make(initialisation.Field)})
		dataModelPtr := &(*dataModel)[len(*dataModel)-1]
		dataModelPtr.Fields[initialisation.Uuid] = &initialisation.DynamicType{}
		dataModelPtr.Fields[initialisation.Uuid].SetData("", initialisation.Uuid)
		for _, e := range model.Fields {
			dataModelPtr.Fields[e.Name] = &initialisation.DynamicType{}
			dataModelPtr.Fields[e.Name].SetData("", initialisation.Datatype(e.Type))
		}
	}
	displayConfiguration(configuration)
	return true
}
