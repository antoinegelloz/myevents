package listener

import (
	"encoding/hex"
	"github.com/agelloz/reach/bookingService/models"
	"github.com/agelloz/reach/bookingService/persistence"
	"github.com/agelloz/reach/contracts"
	"github.com/agelloz/reach/msgqueue"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DBHandler
}

func (p *EventProcessor) ProcessEvents() {
	eventChan, errChan, err := p.EventListener.Listen("event.created", "event.deleted")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case newEvent := <-eventChan:
			p.handleEvent(newEvent)
		case err = <-errChan:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	log.Println("handling new event...")
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("Adding new event from queue ID:%s\n", hex.EncodeToString(e.ID))
		var newID bson.ObjectId
		if !bson.IsObjectIdHex(hex.EncodeToString(e.ID)) {
			log.Printf("Not valid ID|%s|", hex.EncodeToString(e.ID))
			newID = bson.NewObjectId()
		} else {
			newID = bson.ObjectIdHex(hex.EncodeToString(e.ID))
		}
		var newLocation bson.ObjectId
		if !bson.IsObjectIdHex(hex.EncodeToString(e.LocationID)) {
			newLocation = bson.NewObjectId()
		} else {
			newLocation = bson.ObjectIdHex(hex.EncodeToString(e.LocationID))
		}
		_, err := p.Database.AddEvent(models.Event{
			ID:         newID,
			Name:       e.Name,
			LocationID: newLocation,
			Start:      e.Start,
			End:        e.End,
		})
		if err != nil {
			log.Panicf("error while adding event to bookingService database: %s", err)
		}
		log.Printf("event %s added to bookingService database: %+v", hex.EncodeToString(e.ID), e)
	case *contracts.EventDeletedEvent:
		log.Printf("Deleting event from queue ID:%s\n", hex.EncodeToString(e.ID))
		var newID bson.ObjectId
		if !bson.IsObjectIdHex(hex.EncodeToString(e.ID)) {
			log.Printf("Not valid ID |%s|", hex.EncodeToString(e.ID))
			newID = bson.NewObjectId()
		} else {
			newID = bson.ObjectIdHex(hex.EncodeToString(e.ID))
		}
		err := p.Database.DeleteEvent(models.Event{
			ID: newID,
		})
		if err != nil {
			log.Panicf("error while deleting event from bookingService database: %s", err)
		}
		log.Printf("event %s deleted from bookingService database: %+v", hex.EncodeToString(e.ID), e)
	default:
		log.Printf("unknown event: %t", e)
	}
}
