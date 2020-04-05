package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reactapp/eventsService/models"
)

func (eh *eventsServiceHandler) addEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Cannot decode event data", http.StatusInternalServerError)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot add event ID: %s", id), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Added new event ID:%d\n", id)
}
