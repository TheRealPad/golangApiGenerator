package initialisation

import (
	"encoding/json"
	"github.com/google/uuid"
	"strconv"
)

type Datatype string

const (
	Boolean Datatype = "boolean"
	Integer          = "integer"
	Float            = "float"
	Complex          = "complex"
	String           = "string"
	Uuid             = "uuid"
)

type dynamicTypeInterface interface {
	SetData(data string, dataType Datatype)
	GetData() interface{}
	GetDataType() Datatype
	isDataTypeValid(dataType Datatype) bool
}

type DynamicType struct {
	data     interface{}
	dataType Datatype
}

func (d *DynamicType) SetData(data string, dataType Datatype) {
	if !d.isDataTypeValid(dataType) {
		d.data = data
		d.dataType = String
		return
	}
	d.dataType = dataType
	switch dataType {
	case Boolean:
		d.data = data == "true"
	case Integer:
		d.data, _ = strconv.Atoi(data)
	case Float:
		d.data, _ = strconv.ParseFloat(data, 64)
	case Complex:
		d.data, _ = strconv.ParseComplex(data, 64)
	case Uuid:
		d.data, _ = uuid.Parse(data)
	default:
		d.data = data
	}
}

func (d DynamicType) GetData() interface{} {
	return d.data
}

func (d DynamicType) GetDataType() Datatype {
	return d.dataType
}

func (d DynamicType) isDataTypeValid(dataType Datatype) bool {
	switch dataType {
	case Boolean, Integer, Float, Complex, String, Uuid:
		return true
	}
	return false
}

func (d *DynamicType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.data)
}

func InterfaceToDataModel(data map[string]interface{}, dataModel DataModel) DataModel {
	d := &DataModel{}
	d.Fields = make(Field)
	for key, value := range data {
		_, ok := dataModel.Fields[key]
		if ok {
			d.Fields[key] = &DynamicType{}
			switch value.(type) {
			case string:
				d.Fields[key].SetData(value.(string), String)
			case int:
				d.Fields[key].SetData(strconv.Itoa(value.(int)), Integer)
			case float32:
				dataString := strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32)
				d.Fields[key].SetData(dataString, Float)
			case float64:
				dataString := strconv.FormatFloat(value.(float64), 'f', -1, 64)
				d.Fields[key].SetData(dataString, Float)
			case rune:
				d.Fields[key].SetData(strconv.Itoa(int(value.(int32))), Integer)
			case bool:
				dataString := strconv.FormatBool(value.(bool))
				d.Fields[key].SetData(dataString, Boolean)
			}
		}
	}
	return *d
}

type Field map[string]*DynamicType

type DataModel struct {
	Name   string
	Fields Field
}
