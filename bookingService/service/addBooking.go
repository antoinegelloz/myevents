package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/agelloz/reach/bookingService/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (eh *BookingServiceHandler) AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		http.Error(w, fmt.Sprintf("Bad request to add new booking ID: %s", eventID), http.StatusBadRequest)
		return
	}

	var event models.Event
	var err error
	id, err := hex.DecodeString(eventID)
	if err == nil {
		event, err = eh.DBHandler.GetEventByID(id)
	}
	if err != nil {
		http.Error(w, "Cannot find event to book by ID", http.StatusNotFound)
		return
	}
	log.Printf("got event to book by ID %s\n", event.ID)

	newBooking := models.Booking{}
	err = json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		http.Error(w, "Cannot decode booking data", http.StatusInternalServerError)
		return
	}
	id, err = eh.DBHandler.AddBooking(newBooking)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot add new booking ID: %s", id), http.StatusInternalServerError)
		return
	}
	log.Printf("added new booking ID:%s\n", newBooking.ID.String())
}
