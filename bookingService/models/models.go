package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Booking represents an event booking
type Booking struct {
	ID       bson.ObjectId `bson:"_id"`
	EventID  bson.ObjectId
	UserID   bson.ObjectId
	Date     int64
	Quantity uint
}

// Event represents an event
type Event struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string
	LocationID bson.ObjectId
	Start      time.Time
	End        time.Time
}
