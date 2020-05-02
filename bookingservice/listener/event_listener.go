package listener

import (
	"encoding/hex"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"github.com/agelloz/myevents/bookingservice/models"
	"github.com/agelloz/myevents/bookingservice/persistence"
	"github.com/agelloz/myevents/contracts"
	"github.com/streadway/amqp"
)

type Event interface {
	EventName() string
}

func Listen(AMQPMessageBroker string, dh persistence.DBHandler) {
	var err error
	var conn *amqp.Connection
	log.Println("connecting to AMQP message broker...")
	conn, err = amqp.Dial(AMQPMessageBroker)
	for err != nil {
		log.Printf("AMQP connection error: %s\n", err)
		time.Sleep(2000000000)
		conn, err = amqp.Dial(AMQPMessageBroker)
	}
	defer conn.Close()

	var ch *amqp.Channel
	log.Println("opening channel to AMQP message broker...")
	ch, err = conn.Channel()
	for err != nil {
		log.Printf("AMQP channel error: %s\n", err)
		time.Sleep(2000000000)
		ch, err = conn.Channel()
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("events_queue", false, false, false, false, nil)
	if err != nil {
		log.Printf("queue declaring error: %s\n", err)
		return
	}
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("channel consuming error: %s\n", err)
		return
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
}

func HandleEvent(dh persistence.DBHandler, event Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		objID, err := primitive.ObjectIDFromHex(hex.EncodeToString(e.ID))
		if err != nil {
			log.Fatal(err)
		}
		newID := dh.AddEvent(&models.Event{
			ID:        objID,
			Name:      e.Name,
			StartDate: e.StartDate,
			EndDate:   e.EndDate,
			Location: models.Location{
				Name:    e.LocationName,
				Country: e.LocationCountry,
			},
		})
		if newID != objID {
			log.Printf("error while adding event to database: %s", err)
		}
	case *contracts.EventDeletedEvent:
		eventToDelete := dh.GetEventByID(hex.EncodeToString(e.ID))
		dh.DeleteEvent(eventToDelete)
	default:
		log.Printf("unknown event: %t", e)
	}
}
