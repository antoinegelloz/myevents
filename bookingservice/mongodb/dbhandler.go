package mongodb

import (
	"context"
	"time"

	"github.com/agelloz/myevents/bookingservice/models"
	"gopkg.in/mgo.v2/bson"
)

// AddEvent adds an event
func (mgoLayer *DBLayer) AddEvent(e models.Event) (bson.ObjectId, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, e)
	return res.InsertedID.(bson.ObjectId), err
}

// DeleteEvent deletes an event
func (mgoLayer *DBLayer) DeleteEvent(e models.Event) error {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.DeleteOne(ctx, e)
	return err
}

/*
// DeleteAllEvents deletes all events
func (mgoLayer *DBLayer) DeleteAllEvents() error {
	s := mgoLayer.session.Copy()
	defer s.Close()
	_, err := s.DB(DB).C(EVENTS).RemoveAll(nil)
	return err
}
*/

// GetEventByID returns an event
func (mgoLayer *DBLayer) GetEventByID(id []byte) (models.Event, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)
	e := models.Event{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.ObjectId(id)).Decode(&e)
	return e, err
}

/*
// GetEventByName returns an event
func (mgoLayer *DBLayer) GetEventByName(name string) (models.Event, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	var e models.Event
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}
*/

// GetAllEvents returns all available events
func (mgoLayer *DBLayer) GetAllEvents() ([]models.Event, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)
	var events []models.Event
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.D{})
	cursor.Decode(&events)
	return events, err
}

/*
// AddBooking adds a booking
func (mgoLayer *DBLayer) AddBooking(b models.Booking) ([]byte, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	if !b.ID.Valid() {
		b.ID = bson.NewObjectId()
	}
	if !b.UserID.Valid() {
		b.UserID = bson.NewObjectId()
	}
	return []byte(b.ID), s.DB(DB).C(BOOKINGS).Insert(b)
}

// DeleteBooking deletes a booking
func (mgoLayer *DBLayer) DeleteBooking(b models.Booking) error {
	s := mgoLayer.session.Copy()
	defer s.Close()
	return s.DB(DB).C(BOOKINGS).Remove(b)
}

// GetAllBookings returns all bookings
func (mgoLayer *DBLayer) GetAllBookings() ([]models.Booking, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	var bookings []models.Booking
	err := s.DB(DB).C(BOOKINGS).Find(nil).All(&bookings)
	return bookings, err
}

// GetBookingByID returns an booking
func (mgoLayer *DBLayer) GetBookingByID(id []byte) (models.Booking, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	b := models.Booking{}
	err := s.DB(DB).C(BOOKINGS).FindId(bson.ObjectId(id)).One(&b)
	return b, err
}
*/
