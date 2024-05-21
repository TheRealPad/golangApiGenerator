package database

import (
	"github.com/google/uuid"
	"httpServer/src/initialisation"
)

type DatabaseInterface interface {
	Create(data initialisation.DataModel) initialisation.Field
	ReadOne(uuid uuid.UUID, dataModel initialisation.DataModel) (initialisation.Field, error)
	ReadMany(name string) []initialisation.Field
	Update(uuid uuid.UUID, data initialisation.DataModel) initialisation.Field
	Delete(uuid uuid.UUID, name string) bool
}
