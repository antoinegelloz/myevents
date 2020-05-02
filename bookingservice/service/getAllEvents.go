package service

import (
	"encoding/json"
	"net/http"
)

func (eh *BookingServiceHandler) GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events := eh.DBHandler.GetAllEvents()
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err := json.NewEncoder(w).Encode(&events)
	if err != nil {
		http.Error(w, "Cannot encode all events to JSON", http.StatusInternalServerError)
	}
}
