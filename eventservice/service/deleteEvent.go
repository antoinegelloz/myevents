package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	var event models.Event
	var err error
	switch strings.ToLower(nameOrID) {
	case "name":
		event, err = eh.DbHandler.GetEventByName(nameOrIDValue)
		if err != nil {
			http.Error(w, "Cannot get event to delete by name", http.StatusNotFound)
			return
		}
		log.Printf("found event to delete by name %s\n", nameOrIDValue)
	case "id":
		id, err := hex.DecodeString(nameOrIDValue)
		if err == nil {
			event, err = eh.DbHandler.GetEventByID(id)
		}
		if err != nil {
			http.Error(w, "Cannot find event to delete by ID", http.StatusNotFound)
			return
		}
		log.Printf("got event to delete by ID %s\n", event.ID)
	}
	err = eh.DbHandler.DeleteEvent(event)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot delete event ID: %s", event.ID), http.StatusInternalServerError)
		return
	}
	log.Printf("deleted event from database ID:%s\n", event.ID)

	msg := contracts.EventDeletedEvent{
		ID: []byte(event.ID),
	}
	jsonDoc, err := json.Marshal(&msg)
	if nil != err {
		http.Error(w, "error marshal message", http.StatusInternalServerError)
		return
	}
	channel, err := eh.AMQPConnection.Channel()
	if nil != err {
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
	if nil != err {
		http.Error(w, "error sending message", http.StatusInternalServerError)
		return
	}
	log.Print("deletion of event successfully emitted\n")
}
