package service

/*

import (
	"encoding/hex"
	"fmt"
	"github.com/agelloz/myevents/bookingservice/models"
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
	var booking models.Booking
	id, err := hex.DecodeString(eventID)
	if err == nil {
		booking, err = eh.DBHandler.GetBookingByID(id)
	}
	if err != nil {
		http.Error(w, "Cannot find booking to delete by ID", http.StatusNotFound)
		return
	}
	log.Printf("got event to delete a booking by ID %s\n", booking.ID)

	err = eh.DBHandler.DeleteBooking(booking)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot delete booking ID: %s", booking.ID), http.StatusInternalServerError)
		return
	}
	log.Printf("deleted booking from database ID:%s\n", booking.ID)
}
*/
