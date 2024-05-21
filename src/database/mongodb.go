package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"httpServer/src/initialisation"
)

type MongoDB struct {
	Name string
	Url  string
}

func (m MongoDB) Create(data initialisation.DataModel) initialisation.Field {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Url))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection(data.Name)

	document := map[string]interface{}{}
	for key, field := range data.Fields {
		if key == "uuid" && field.GetDataType() == initialisation.Uuid {
			document[key] = field.GetData().(uuid.UUID).String()
		} else {
			document[key] = field.GetData()
		}
	}

	_, err = coll.InsertOne(context.TODO(), document)
	if err != nil {
		panic(err)
	}
	return data.Fields
}

func (m MongoDB) ReadOne(uuid uuid.UUID, dataModel initialisation.DataModel) (initialisation.Field, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Url))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection(dataModel.Name)
	var result map[string]interface{}
	err = coll.FindOne(context.TODO(), bson.D{{"uuid", uuid.String()}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No data was found with the uuid %s\n", uuid)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	d := initialisation.InterfaceToDataModel(result, dataModel)
	return d.Fields, nil
}

func (m MongoDB) ReadMany(name string) []initialisation.Field {
	var lst []initialisation.Field
	d := initialisation.DataModel{}

	lst = append(lst, d.Fields)
	lst = append(lst, d.Fields)
	lst = append(lst, d.Fields)
	return lst
}

func (m MongoDB) Update(uuid uuid.UUID, data initialisation.DataModel) initialisation.Field {
	data.Fields[initialisation.Uuid].SetData(uuid.String(), initialisation.Uuid)
	return data.Fields
}

func (m MongoDB) Delete(uuid uuid.UUID, name string) bool {
	return true
}
