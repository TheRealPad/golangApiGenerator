package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"httpServer/src/initialisation"
	"log"
)

type MongoDB struct {
	Name string
	Url  string
}

func (m MongoDB) Create(data initialisation.DataModel) (initialisation.Field, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Url))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(m.Name).Collection(data.Name)

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
		return nil, err
	}
	return data.Fields, nil
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

	coll := client.Database(m.Name).Collection(dataModel.Name)
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

func (m MongoDB) ReadMany(dataModel initialisation.DataModel) ([]initialisation.Field, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Url))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(m.Name).Collection(dataModel.Name)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var results []map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	var lst []initialisation.Field
	for _, result := range results {
		lst = append(lst, initialisation.InterfaceToDataModel(result, dataModel).Fields)
	}
	return lst, nil
}

func (m MongoDB) Update(uuid uuid.UUID, data initialisation.DataModel) (initialisation.Field, error) {
	data.Fields[initialisation.Uuid].SetData(uuid.String(), initialisation.Uuid)
	return data.Fields, nil
}

func (m MongoDB) Delete(uuid uuid.UUID, name string) (bool, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Url))
	if err != nil {
		return false, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(m.Name).Collection(name)
	_, errorDb := coll.DeleteOne(context.TODO(), bson.D{{"uuid", uuid.String()}})
	if errorDb != nil {
		return false, err
	}
	return true, nil
}
