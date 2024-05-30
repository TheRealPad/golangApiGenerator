package database

import (
	"github.com/google/uuid"
	"httpServer/src/initialisation"
)

type DatabaseInterface interface {
	Create(data initialisation.DataModel) (initialisation.Field, error)
	ReadOne(uuid uuid.UUID, dataModel initialisation.DataModel) (initialisation.Field, error)
	ReadMany(dataModel initialisation.DataModel) ([]initialisation.Field, error)
	Update(uuid uuid.UUID, data initialisation.DataModel) (initialisation.Field, error)
	Delete(uuid uuid.UUID, name string) (bool, error)
}
