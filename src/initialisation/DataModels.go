package initialisation

import "strconv"

type Datatype string

const (
	Boolean Datatype = "boolean"
	Integer          = "integer"
	Float            = "float"
	Complex          = "complex"
	String           = "string"
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
	case Boolean, Integer, Float, Complex, String:
		return true
	}
	return false
}

type Field map[string]*DynamicType

type DataModel struct {
	Name   string
	Fields Field
}
