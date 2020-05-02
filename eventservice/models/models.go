package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
