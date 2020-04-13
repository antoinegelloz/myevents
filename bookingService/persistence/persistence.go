package persistence

import (
	"github.com/agelloz/reach/bookingService/models"
	"github.com/agelloz/reach/bookingService/mongodb"
)

// DBType is type
type DBType string

const (
	// MONGODB is
	MONGODB DBType = "mongodb"
)

// DBHandler is used to communicate with the database
type DBHandler interface {
	AddEvent(models.Event) ([]byte, error)
	DeleteEvent(models.Event) error
	GetEventByID([]byte) (models.Event, error)
	GetEventByName(string) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
	AddBooking(models.Booking) ([]byte, error)
	DeleteBooking(models.Booking) error
	GetBookingByID([]byte) (models.Booking, error)
	GetAllBookings() ([]models.Booking, error)
}

// NewPersistenceLayer is
func NewPersistenceLayer(options DBType, connection string) (DBHandler, error) {
	switch options {
	case MONGODB:
		return mongodb.NewDBLayer(connection)
	}
	return nil, nil
}
