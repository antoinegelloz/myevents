package service

import (
	"encoding/json"
	"net/http"
)

func (eh *EventsServiceHandler) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events := eh.DbHandler.GetAllEvents()
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err := json.NewEncoder(w).Encode(&events)
	if err != nil {
		http.Error(w, "Cannot encode all events to JSON", http.StatusInternalServerError)
	}
}
