package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (eh *BookingServiceHandler) DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingID, ok := vars["bookingID"]
	if !ok {
		http.Error(w, "Bad request (nameOrID)", http.StatusBadRequest)
		return
	}
	booking := eh.DBHandler.GetBookingByID(bookingID)
	if booking == nil {
		log.Println("DeleteBookingHandler: unknown booking ID")
		http.Error(w, "Unknown booking to delete", http.StatusInternalServerError)
		return
	}
	eh.DBHandler.DeleteBooking(booking)
}
