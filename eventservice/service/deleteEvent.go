package service

import (
	"encoding/hex"
	"encoding/json"
	"github.com/agelloz/myevents/contracts"
	"github.com/agelloz/myevents/eventservice/models"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (eh *EventsServiceHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameOrID, ok := vars["nameOrID"]
	if !ok {
		http.Error(w, "Bad request (nameOrID)", http.StatusBadRequest)
		return
	}
	nameOrIDValue, ok := vars["nameOrIDValue"]
	if !ok {
		http.Error(w, "Bad request (nameOrIDValue)", http.StatusBadRequest)
		return
	}
	var event *models.Event
	switch strings.ToLower(nameOrID) {
	case "name":
		event = eh.DbHandler.GetEventByName(nameOrIDValue)
		if event == nil {
			http.Error(w, "Cannot get event to delete by name", http.StatusNotFound)
			return
		}
		log.Printf("found event to delete by name %s\n", nameOrIDValue)
	case "id":
		id, err := hex.DecodeString(nameOrIDValue)
		if err != nil {
			http.Error(w, "Cannot get event to delete by ID", http.StatusNotFound)
			return
		}
		event = eh.DbHandler.GetEventByID(hex.EncodeToString(id))
		if event == nil {
			http.Error(w, "Cannot find event to delete by ID", http.StatusNotFound)
			return
		}
	}
	eh.DbHandler.DeleteEvent(event)

	msg := contracts.EventDeletedEvent{
		ID: event.ID,
	}
	jsonDoc, err := json.Marshal(&msg)
	if err != nil {
		http.Error(w, "error marshal message", http.StatusInternalServerError)
		return
	}
	channel, err := eh.AMQPConnection.Channel()
	if err != nil {
		http.Error(w, "error opening channel", http.StatusInternalServerError)
		return
	}
	defer channel.Close()
	q, err := channel.QueueDeclare("events_queue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	message := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": "event.deleted"},
		Body:        jsonDoc,
		ContentType: "application/json",
	}
	err = channel.Publish("", q.Name, false, false, message)
	if err != nil {
		http.Error(w, "error sending message", http.StatusInternalServerError)
		return
	}
	log.Printf("deletion of event successfully emitted with ID: %+v\n", event.ID)
}
