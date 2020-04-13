package mongodb

import (
	"github.com/agelloz/reach/bookingService/models"
	"gopkg.in/mgo.v2/bson"
)

// AddEvent adds an event
func (mgoLayer *DBLayer) AddEvent(e models.Event) ([]byte, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

// DeleteEvent deletes an event
func (mgoLayer *DBLayer) DeleteEvent(e models.Event) error {
	s := mgoLayer.session.Copy()
	defer s.Close()
	return s.DB(DB).C(EVENTS).Remove(e)
}

// GetEventByID returns an event
func (mgoLayer *DBLayer) GetEventByID(id []byte) (models.Event, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	e := models.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

// GetEventByName returns an event
func (mgoLayer *DBLayer) GetEventByName(name string) (models.Event, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	var e models.Event
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

// GetAllEvents returns all available events
func (mgoLayer *DBLayer) GetAllEvents() ([]models.Event, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	var events []models.Event
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

// AddBooking adds a booking
func (mgoLayer *DBLayer) AddBooking(b models.Booking) ([]byte, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	if !b.ID.Valid() {
		b.ID = bson.NewObjectId()
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
