package models

type Configuration struct {
	Port   int         `json:"port"`
	Name   string      `json:"name"`
	Models []DataModel `json:"dataModels"`
	Db     Database    `json:"database"`
}
