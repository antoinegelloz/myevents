package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Booking represents an event booking
type Booking struct {
	ID        primitive.ObjectID `bson:"_id"`
	EventID   primitive.ObjectID `bson:"event_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	UserEmail string             `json:"user_email"`
	Date      time.Time          `json:"booking_date"`
	Quantity  uint               `json:"booking_quantity"`
}

// Event represents an event
type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
	Location  Location           `json:"location"`
}

// Location represents a location
type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
