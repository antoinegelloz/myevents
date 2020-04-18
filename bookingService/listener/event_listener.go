package listener

import (
	"encoding/hex"
	"encoding/json"
	"github.com/agelloz/reach/bookingService/models"
	"github.com/agelloz/reach/bookingService/persistence"
	"github.com/agelloz/reach/contracts"
	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Event interface {
	EventName() string
}

type AMQPEventListener struct {
	Connection *amqp.Connection
	Queue      string
	Channel    *amqp.Channel
}

func Listen(conn *amqp.Connection, dh persistence.DBHandler) error {
	ch, err := conn.Channel()
	if err != nil {
		log.Println("channel error")
		return err
	}
	defer ch.Close()
	messages, err := ch.Consume("events_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Println("channel consuming error")
		return err
	}
	log.Println("listening to events...")
	forever := make(chan bool)
	go func() {
		for msg := range messages {
			// Map message to actual event struct
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				log.Println("msg did not contain x-event-name header")
				err := msg.Nack(false, false)
				if err != nil {
					log.Printf("nack error: %s\n", err)
				}
				continue
			}
			eventName, ok := rawEventName.(string)
			if !ok {
				log.Printf("x-event-name header is not string, but %t\n", rawEventName)
				err := msg.Nack(false, false)
				if err != nil {
					log.Printf("nack error: %s\n", err)
				}
				continue
			}
			var event Event
			switch eventName {
			case "event.created":
				event = new(contracts.EventCreatedEvent)
			case "event.deleted":
				event = new(contracts.EventDeletedEvent)
			default:
				log.Printf("event type %s is unknown\n", eventName)
				continue
			}
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				log.Printf("unmarshal error: %s\n", err)
				continue
			}
			HandleEvent(dh, event)
		}
	}()
	<-forever
	return nil
}

func HandleEvent(dh persistence.DBHandler, event Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
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
		_, err := dh.AddEvent(models.Event{
			ID:         newID,
			Name:       e.Name,
			LocationID: newLocation,
			Start:      e.Start,
			End:        e.End,
		})
		if err != nil {
			log.Printf("error while adding event to database: %s", err)
		} else {
			log.Printf("added event %s to database: %+v", hex.EncodeToString(e.ID), e)
		}
	case *contracts.EventDeletedEvent:
		if !bson.IsObjectIdHex(hex.EncodeToString(e.ID)) {
			log.Printf("error while deleting event from database: invalid ID |%s|", hex.EncodeToString(e.ID))
		} else {
			foundEvent, err := dh.GetEventByID(e.ID)
			if err != nil {
				log.Printf("error while deleting event from database: %s", err)
			} else {
				err := dh.DeleteEvent(foundEvent)
				if err != nil {
					log.Printf("error while deleting event from database: %s", err)
				} else {
					log.Printf("deleted event %s from database: %+v", hex.EncodeToString(e.ID), e)
				}
			}
		}
	default:
		log.Printf("unknown event: %t", e)
	}
}
