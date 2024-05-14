package models

import (
	"httpServer/src/models"
	"reflect"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	log := models.Log{
		Method:  "GET",
		Url:     "/example",
		Address: "127.0.0.1",
		Time:    time.Now(),
	}
	expectedType := reflect.TypeOf(log)
	if expectedType.Field(0).Type.Kind() != reflect.String {
		t.Errorf("Method field type is not string")
	}
	if expectedType.Field(1).Type.Kind() != reflect.String {
		t.Errorf("Url field type is not string")
	}
	if expectedType.Field(2).Type.Kind() != reflect.String {
		t.Errorf("Address field type is not string")
	}
	if expectedType.Field(3).Type.Kind() != reflect.Struct {
		t.Errorf("Time field type is not time.Time")
	}
}
