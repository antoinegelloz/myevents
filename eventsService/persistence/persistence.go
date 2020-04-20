package persistence

import (
	"github.com/agelloz/reach/eventsService/models"
	"github.com/agelloz/reach/eventsService/mongodb"
)

// DBTYPE is type
type DBTYPE string

const (
	// MONGODB is
	MONGODB DBTYPE = "mongodb"
)

// DBHandler is used to communicate with the database
type DBHandler interface {
	AddEvent(models.Event) ([]byte, error)
	DeleteEvent(models.Event) error
	DeleteAllEvents() error
	GetEventByID([]byte) (models.Event, error)
	GetEventByName(string) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
}

// NewPersistenceLayer is
func NewPersistenceLayer(options DBTYPE, connection string) (DBHandler, error) {
	switch options {
	case MONGODB:
		return mongodb.NewDBLayer(connection)
	}
	return nil, nil
}
