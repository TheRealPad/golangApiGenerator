package initialisation

import (
	"httpServer/src/initialisation"
	"testing"
)

func TestNewDataType(t *testing.T) {
	var data initialisation.DynamicType

	data.SetData("Hello there", initialisation.String)
	if data.GetDataType() != initialisation.String {
		t.Errorf("Bad data type for data got: %s", data.GetDataType())
	}
	if data.GetData() != "Hello there" {
		t.Errorf("Bad data value for data got: %s", data.GetData())
	}
	data.SetData("10", initialisation.Integer)
	if data.GetDataType() != initialisation.Integer {
		t.Errorf("Bad data type for data got: %s", data.GetDataType())
	}
	if data.GetData() != 10 {
		t.Errorf("Bad data value for data got: %s", data.GetData())
	}
}

func TestFields(t *testing.T) {
	fields := make(initialisation.Field)

	fields["name"] = &initialisation.DynamicType{}
	fields["name"].SetData("John", initialisation.String)
	fields["age"] = &initialisation.DynamicType{}
	fields["age"].SetData("10", initialisation.Integer)
	if fields["name"].GetData() != "John" || fields["name"].GetDataType() != initialisation.String {
		t.Errorf("Error with name field, received %s and %s", fields["name"].GetData(), fields["name"].GetDataType())
	}
	if fields["age"].GetData() != 10 || fields["age"].GetDataType() != initialisation.Integer {
		t.Errorf("Error with name field, received %s and %s", fields["age"].GetData(), fields["age"].GetDataType())
	}
}

func TestDataLModel(t *testing.T) {
	dataModel := initialisation.DataModel{Name: "User", Fields: make(initialisation.Field)}

	dataModel.Fields["name"] = &initialisation.DynamicType{}
	dataModel.Fields["name"].SetData("John", initialisation.String)
	dataModel.Fields["age"] = &initialisation.DynamicType{}
	dataModel.Fields["age"].SetData("10", initialisation.Integer)
	if dataModel.Fields["name"].GetData() != "John" || dataModel.Fields["name"].GetDataType() != initialisation.String {
		t.Errorf("Error with name field, received %s and %s", dataModel.Fields["name"].GetData(), dataModel.Fields["name"].GetDataType())
	}
	if dataModel.Fields["age"].GetData() != 10 || dataModel.Fields["age"].GetDataType() != initialisation.Integer {
		t.Errorf("Error with name field, received %s and %s", dataModel.Fields["age"].GetData(), dataModel.Fields["age"].GetDataType())
	}
}
