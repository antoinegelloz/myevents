package persistence

import (
	"github.com/agelloz/myevents/eventservice/models"
	"github.com/agelloz/myevents/eventservice/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DBType is type
type DBType string

const (
	// MONGODB is
	MONGODB DBType = "mongodb"
)

// DBHandler is used to communicate with the database
type DBHandler interface {
	AddEvent(e *models.Event) primitive.ObjectID
	DeleteEvent(e *models.Event)
	DeleteAllEvents()
	GetEventByID(ID string) *models.Event
	GetEventByName(name string) *models.Event
	GetAllEvents() []*models.Event
}

// NewPersistenceLayer is
func NewPersistenceLayer(options DBType, uri string) DBHandler {
	switch options {
	case MONGODB:
		return mongodb.NewDBLayer(uri)
	}
	return nil
}
