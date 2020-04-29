package mongodb

import (
	"github.com/agelloz/myevents/eventservice/models"

	"gopkg.in/mgo.v2/bson"
)

// AddEvent adds an event
func (mgoLayer *DBLayer) AddEvent(e models.Event) ([]byte, error) {
	s := mgoLayer.session.Copy()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	// let's assume the method below checks if the ID is valid for the location object of the event
	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
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

// DeleteAllEvents deletes all events
func (mgoLayer *DBLayer) DeleteAllEvents() error {
	s := mgoLayer.session.Copy()
	defer s.Close()
	_, err := s.DB(DB).C(EVENTS).RemoveAll(nil)
	return err
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
