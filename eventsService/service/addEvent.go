package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/agelloz/reach/contracts"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"time"

	"github.com/agelloz/reach/eventsService/models"
)

func (eh *EventsServiceHandler) AddEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Cannot decode event data", http.StatusInternalServerError)
		return
	}
	id, err := eh.DbHandler.AddEvent(event)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot add event ID: %s", hex.EncodeToString(id)), http.StatusInternalServerError)
		return
	}
	log.Printf("added new event to database ID:%s\n", hex.EncodeToString(id))

	msg := contracts.EventCreatedEvent{
		ID:         id,
		Name:       event.Name,
		LocationID: []byte(event.Location.ID),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}

	jsonDoc, err := json.Marshal(&msg)
	if nil != err {
		http.Error(w, "error marshal message", http.StatusInternalServerError)
		return
	}
	channel, err := eh.AMQPConnection.Channel()
	if nil != err {
		http.Error(w, "error channel", http.StatusInternalServerError)
		return
	}
	defer channel.Close()
	message := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": "event.created"},
		Body:        jsonDoc,
		ContentType: "application/json",
	}
	err = channel.Publish("", "events_queue", false, false, message)
	if nil != err {
		http.Error(w, "error sending message", http.StatusInternalServerError)
		return
	}
	log.Printf("creation of event successfully emitted with ID:%s\n", hex.EncodeToString(id))
}
