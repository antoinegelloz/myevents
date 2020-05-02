package service

import (
	"encoding/json"
	"net/http"
)

func (eh *BookingServiceHandler) GetAllBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings := eh.DBHandler.GetAllBookings()
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err := json.NewEncoder(w).Encode(&bookings)
	if err != nil {
		http.Error(w, "Cannot encode all bookings to JSON", http.StatusInternalServerError)
	}
}
