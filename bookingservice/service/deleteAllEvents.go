package service

import (
	"log"
	"net/http"
)

func (eh *BookingServiceHandler) DeleteAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	eh.DBHandler.DeleteAllEvents()
	log.Println("deleted all events")
}
