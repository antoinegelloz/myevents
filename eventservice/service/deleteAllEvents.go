package service

import (
	"log"
	"net/http"
)

func (eh *EventsServiceHandler) DeleteAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	err := eh.DbHandler.DeleteAllEvents()
	if err != nil {
		http.Error(w, "Cannot delete all events", http.StatusInternalServerError)
		return
	}
	log.Println("deleted all events")
}
