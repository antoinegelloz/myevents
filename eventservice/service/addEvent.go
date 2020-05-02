package service

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"

	"github.com/agelloz/myevents/contracts"
	"github.com/streadway/amqp"

	"github.com/agelloz/myevents/eventservice/models"
)

func (eh *EventsServiceHandler) AddEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Cannot decode event data", http.StatusInternalServerError)
		return
	}

	event.ID = primitive.NewObjectID()

	// Default values
	if event.Name == "" {
		event.Name = "event"
	}
	if event.StartDate.IsZero() {
		event.StartDate = time.Now()
	}
	if event.EndDate.IsZero() {
		event.EndDate = time.Now()
	}
	if event.Location.Country == "" {
		event.Location.Country = "France"
	}
	if event.Location.Name == "" {
		event.Location.Name = "Paris"
	}

	res := eh.DbHandler.AddEvent(&event)
	log.Print(res)

	msg := contracts.EventCreatedEvent{
		ID:              [12]byte(res),
		Name:            event.Name,
		StartDate:       event.StartDate,
		EndDate:         event.EndDate,
		LocationName:    event.Location.Name,
		LocationCountry: event.Location.Country,
	}

	jsonDoc, err := json.Marshal(&msg)
	if err != nil {
		log.Printf("AddEventHandler: cannot marshal message: %+v\n", msg)
		http.Error(w, "Cannot add new event: error marshal message", http.StatusInternalServerError)
		return
	}
	channel, err := eh.AMQPConnection.Channel()
	if nil != err {
		log.Println("AddEventHandler: cannot open AMQP channel")
		http.Error(w, "error opening channel", http.StatusInternalServerError)
		return
	}
	defer channel.Close()
	q, err := channel.QueueDeclare("events_queue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	message := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": "event.created"},
		Body:        jsonDoc,
		ContentType: "application/json",
	}
	err = channel.Publish("", q.Name, false, false, message)
	if nil != err {
		http.Error(w, "error sending message", http.StatusInternalServerError)
		return
	}
	log.Printf("creation of event successfully emitted with ID: %+v\n", res)
}
