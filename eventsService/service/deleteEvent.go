package service

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/agelloz/reach/eventsService/models"

	"github.com/gorilla/mux"
)

func (eh *EventsServiceHandler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
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
		event, err = eh.dbHandler.GetEventByName(nameOrIDValue)
		if err != nil {
			http.Error(w, "Cannot get event to delete by name", http.StatusNotFound)
			return
		}
		fmt.Printf("Found event to delete by name %s\n", nameOrIDValue)
	case "id":
		id, err := hex.DecodeString(nameOrIDValue)
		if err == nil {
			event, err = eh.dbHandler.GetEventByID(id)
		}
		if err != nil {
			http.Error(w, "Cannot find event to delete by ID", http.StatusNotFound)
			return
		}
		fmt.Printf("Got event to delete by ID %s\n", event.ID)
	}
	err = eh.dbHandler.DeleteEvent(event)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot add event ID: %s", event.ID), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Delete event ID:%s\n", event.ID)
}
