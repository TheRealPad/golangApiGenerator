package initialisation

import "fmt"

type Settings struct {
	Port   int        `json:"port"`
	Name   string     `json:"name"`
	Models DataModels `json:"dataModels"`
	Db     Database   `json:"database"`
}

type DataModels struct {
	DataModels []DataModel `json:"dataModels"`
}

type DataModel struct {
	Name     string `json:"name"`
	Create   bool   `json:"create"`
	ReadOne  bool   `json:"readOne"`
	ReadMany bool   `json:"readMany"`
	Update   bool   `json:"update"`
	Delete   bool   `json:"delete"`
	Fields   Fields `json:"fields"`
}

type Fields struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Value string
}

type Database struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}

type JsonHandlerInterface interface {
	ReadFile()
	GetFieldSingle()
	GetFieldArray()
}

type JsonHandler struct {
	File string
}

func (j JsonHandler) ReadFile() bool {
	fmt.Println(j.File)
	return true
}
