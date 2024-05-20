package database

import (
	"github.com/google/uuid"
	"httpServer/src/initialisation"
)

type MongoDB struct {
}

func (m MongoDB) Create(data initialisation.DataModel) initialisation.Field {
	return data.Fields
}

func (m MongoDB) ReadOne(uuid uuid.UUID, name string) initialisation.Field {
	d := initialisation.DataModel{}
	d.Fields[initialisation.Uuid].SetData(uuid.String(), initialisation.Uuid)
	return d.Fields
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
