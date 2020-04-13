package listener

import (
	"fmt"
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
	log.Println("Listening to events...")

	eventChan, errChan, err := p.EventListener.Listen("event.created", "event.deleted")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case evt := <-eventChan:
			p.handleEvent(evt)
		case err = <-errChan:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		newID := bson.ObjectIdHex(e.ID)
		_, err := p.Database.AddEvent(models.Event{
			ID:         newID,
			Name:       e.Name,
			LocationID: bson.ObjectId(e.LocationID),
			Start:      e.Start,
			End:        e.End,
		})
		if err != nil {
			panic(fmt.Errorf("error while adding event to bookingService database: %s", err))
		}
		log.Printf("event %s added to bookingService database: %s", e.ID, e)
	case *contracts.EventDeletedEvent:
		toDeleteID := bson.ObjectIdHex(e.ID)
		err := p.Database.DeleteEvent(models.Event{
			ID: toDeleteID,
		})
		if err != nil {
			panic(fmt.Errorf("error while deleting event from bookingService database: %s", err))
		}
		log.Printf("event %s deleted from bookingService database: %s", e.ID, e)
	default:
		log.Printf("unknown event: %t", e)
	}
}
