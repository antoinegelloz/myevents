package service

import (
	"encoding/hex"
	"fmt"
	"github.com/agelloz/reach/contracts"
	"github.com/agelloz/reach/eventsService/models"
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
	err = eh.EventEmitter.Emit(&msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot emit deletion of event ID: %s",
			hex.EncodeToString(msg.ID)), http.StatusInternalServerError)
		return
	}
	log.Print("deletion of event successfully emitted\n")
}
