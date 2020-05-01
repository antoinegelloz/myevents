package service

/*

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/agelloz/myevents/bookingservice/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (eh *BookingServiceHandler) AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		log.Printf("bad request to add new booking for event ID: %s\n", eventID)
		http.Error(w, fmt.Sprintf("Bad request to add new booking for event ID: %s", eventID), http.StatusBadRequest)
		return
	}

	var event models.Event
	var err error
	id, err := hex.DecodeString(eventID)
	if err == nil {
		event, err = eh.DBHandler.GetEventByID(id)
	}
	if err != nil {
		log.Println("cannot find event to book by ID")
		http.Error(w, "cannot find event to book by ID", http.StatusNotFound)
		return
	}
	log.Printf("got event to book by ID %s\n", event.ID)
	newBooking := models.Booking{}
	err = json.NewDecoder(r.Body).Decode(&newBooking)
	if err != nil {
		log.Println("cannot decode booking data")
		http.Error(w, "cannot decode booking data", http.StatusInternalServerError)
		return
	}
	newBooking.EventID = bson.ObjectId(id)
	newBooking.Date = time.Now()
	id, err = eh.DBHandler.AddBooking(newBooking)
	if err != nil {
		log.Printf("cannot add new booking for event ID: %s\n", newBooking.EventID.Hex())
		http.Error(w, fmt.Sprintf("cannot add new booking for event ID: %s", newBooking.EventID.Hex()), http.StatusInternalServerError)
		return
	}
	log.Printf("added new booking ID:%s for event ID:%s quantity:%d\n", bson.ObjectId(id).Hex(), newBooking.EventID.Hex(), newBooking.Quantity)

	auth := smtp.PlainAuth("", eh.SMTPUsername, eh.SMTPPassword, eh.SMTPHost)
	to := []string{newBooking.UserEmail}
	msg := []byte("To: " + newBooking.UserEmail + "\r\n" +
		"Subject: See you soon at " + event.Name + "!\r\n" +
		"\r\n" +
		"Your booking has been confirmed. Congratulations!\n\nBooking ID: " + bson.ObjectId(id).Hex() + "\r\n")
	log.Printf("sending mail to:%s message:%s\n", to, msg)
	err = smtp.SendMail(eh.SMTPAddr, auth, eh.SMTPUsername, to, msg)
	if err != nil {
		log.Printf("SMTP error: %s\n", err)
		http.Error(w, fmt.Sprintf("cannot send email for event ID: %s", newBooking.EventID.Hex()), http.StatusInternalServerError)
		return
	}
}
*/
