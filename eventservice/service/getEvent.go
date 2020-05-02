package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agelloz/myevents/eventservice/models"

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
	var event *models.Event
	switch strings.ToLower(nameOrID) {
	case "name":
		event = eh.DbHandler.GetEventByName(nameOrIDValue)
		if event == nil {
			http.Error(w, "Cannot get event by name", http.StatusNotFound)
			return
		}
	case "id":
		event = eh.DbHandler.GetEventByID(nameOrIDValue)
		if event == nil {
			http.Error(w, "Cannot get event by ID", http.StatusNotFound)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err := json.NewEncoder(w).Encode(&event)
	if err != nil {
		http.Error(w, "Cannot encode events to JSON", http.StatusInternalServerError)
	}
}
