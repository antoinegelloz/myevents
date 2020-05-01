package persistence

import (
	"github.com/agelloz/myevents/bookingservice/models"
	"github.com/agelloz/myevents/bookingservice/mongodb"
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
	GetEventByID(ID string) *models.Event
	GetEventByName(name string) *models.Event
	DeleteAllEvents()
	GetAllEvents() []*models.Event
	AddBooking(b *models.Booking) primitive.ObjectID
	DeleteBooking(b *models.Booking)
	GetBookingByID(ID string) *models.Booking
	GetAllBookings() []*models.Booking
}

// NewPersistenceLayer is
func NewPersistenceLayer(options DBType, uri string) DBHandler {
	switch options {
	case MONGODB:
		return mongodb.NewDBLayer(uri)
	}
	return nil
}
