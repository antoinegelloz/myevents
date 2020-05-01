package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/agelloz/myevents/bookingservice/models"
	"github.com/gorilla/mux"
)

func (eh *BookingServiceHandler) AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		log.Printf("bad request to add new booking for event ID: %s\n", eventID)
		http.Error(w, fmt.Sprintf("Bad request to add new booking for event ID: %s", eventID), http.StatusBadRequest)
		return
	}

	event := eh.DBHandler.GetEventByID(eventID)
	log.Printf("got event to book by ID %s\n", event.ID)
	newBooking := models.Booking{}
	err := json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		log.Println("cannot decode booking data")
		http.Error(w, "cannot decode booking data", http.StatusInternalServerError)
		return
	}
	newBooking.Date = time.Now()
	objID := eh.DBHandler.AddBooking(&newBooking)
	log.Printf("added new booking ID:%s for event ID:%s quantity:%d\n", objID.Hex(), newBooking.EventID.Hex(), newBooking.Quantity)

	auth := smtp.PlainAuth("", eh.SMTPUsername, eh.SMTPPassword, eh.SMTPHost)
	to := []string{newBooking.UserEmail}
	msg := []byte("To: " + newBooking.UserEmail + "\r\n" +
		"Subject: See you soon at " + event.Name + "!\r\n" +
		"\r\n" +
		"Your booking has been confirmed. Congratulations!\n\nBooking ID: " + objID.Hex() + "\r\n")
	log.Printf("sending mail to:%s message:%s\n", to, msg)
	err = smtp.SendMail(eh.SMTPAddr, auth, eh.SMTPUsername, to, msg)
	if err != nil {
		log.Printf("SMTP error: %s\n", err)
		http.Error(w, fmt.Sprintf("cannot send email for event ID: %s", newBooking.EventID.Hex()), http.StatusInternalServerError)
		return
	}
}
