package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/agelloz/reach/eventsService/models"

	"github.com/gorilla/mux"
)

func (eh *EventsServiceHandler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, "Cannot get event by name", http.StatusNotFound)
			return
		}
		fmt.Printf("Got event by name %s\n", nameOrIDValue)
	case "id":
		id, err := hex.DecodeString(nameOrIDValue)
		if err == nil {
			event, err = eh.DbHandler.GetEventByID(id)
		}
		if err != nil {
			http.Error(w, "Cannot get event by ID", http.StatusNotFound)
			return
		}
		fmt.Printf("Got event by ID %s\n", event.ID)
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&event)
	if err != nil {
		http.Error(w, "Cannot encode events to JSON", http.StatusInternalServerError)
	}
}
