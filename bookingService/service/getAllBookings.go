package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (eh *BookingServiceHandler) GetAllBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := eh.DBHandler.GetAllBookings()
	if err != nil {
		http.Error(w, "Cannot get all bookings", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&bookings)
	if err != nil {
		http.Error(w, "Cannot encode all bookings to JSON", http.StatusInternalServerError)
	}
	fmt.Println("Got all bookings")
}
