package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (eh *BookingServiceHandler) DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		http.Error(w, "Bad request (nameOrID)", http.StatusBadRequest)
		return
	}
	booking := eh.DBHandler.GetBookingByID(eventID)
	log.Printf("got event to delete a booking by ID %s\n", booking.ID)

	eh.DBHandler.DeleteBooking(booking)
	log.Printf("deleted booking from database ID:%s\n", booking.ID)
}
