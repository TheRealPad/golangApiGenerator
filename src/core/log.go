package core

import (
	"fmt"
	"httpServer/src/initialisation"
	"httpServer/src/models"
)

func displayConfiguration(configuration *models.Configuration) {
	fmt.Println("CONFIGURATION:\n")
	fmt.Println("port:", configuration.Port)
	fmt.Println("name:", configuration.Name)
	fmt.Println("isSecure:", configuration.IsSecure)
	fmt.Println("Database:")
	fmt.Println("\turl:", configuration.Db.Url)
	fmt.Println("\tname:", configuration.Db.Name)
	fmt.Println("data models:")
	fmt.Println("total:", len(configuration.Models))
	for _, model := range configuration.Models {
		fmt.Println("\tname:", model.Name)
		fmt.Print("\tfields:", len(model.Fields), " ")
		for _, e := range model.Fields {
			fmt.Print(e.Name + " - " + e.Type + " ")
		}
		fmt.Println()
		fmt.Println("\tcreate:", model.Create)
		fmt.Println("\tread one:", model.ReadOne)
		fmt.Println("\tread many:", model.ReadMany)
		fmt.Println("\tupdate:", model.Update)
		fmt.Println("\tdelete:", model.Delete)
		fmt.Println()
	}
}

func displayDataTypes(dataModel *[]initialisation.DataModel) {
	fmt.Println("DATA TYPES:\n")
	for _, elem := range *dataModel {
		fmt.Println(elem.Name)
		for k, f := range elem.Fields {
			fmt.Println("\t", k, ":", f.GetDataType())
		}
	}
}
