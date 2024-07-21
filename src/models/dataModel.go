package models

type DataModel struct {
	Name     string  `json:"name"`
	Create   bool    `json:"create"`
	ReadOne  bool    `json:"readOne"`
	ReadMany bool    `json:"readMany"`
	Update   bool    `json:"update"`
	Delete   bool    `json:"delete"`
	Patch    bool    `json:"patch"`
	Fields   []Field `json:"fields"`
}
