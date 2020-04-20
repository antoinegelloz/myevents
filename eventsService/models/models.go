package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Event represents an event
type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `json:"name"`
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	Location  Location      `json:"location"`
}

// Location represents a location
type Location struct {
	ID      bson.ObjectId `bson:"location_id"`
	Name    string        `json:"name"`
	Country string        `json:"country"`
}
