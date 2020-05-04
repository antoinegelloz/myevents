package service

import (
	"github.com/agelloz/myevents/bookingservice/configuration"
	"github.com/agelloz/myevents/bookingservice/listener"
	"github.com/agelloz/myevents/bookingservice/persistence"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type BookingServiceHandler struct {
	DBHandler    persistence.DBHandler
	Endpoint     string
	TLSEndpoint  string
	SMTPUsername string
	SMTPPassword string
	SMTPHost     string
	SMTPAddr     string
}

// ServeAPI is
func ServeAPI() (chan error, chan error) {
	conf, err := configuration.ExtractConfiguration()
	if err != nil {
		panic(err)
	}

	var dh persistence.DBHandler
	log.Println("connecting to database...")
	dh = persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	for dh == nil {
		log.Printf("database connection error: %s\n", err)
		time.Sleep(4000000000)
		dh = persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	}
	log.Println("connected to database")

	go listener.Listen(conf.AMQPMessageBroker, dh)

	eh := &BookingServiceHandler{
		DBHandler:    dh,
		Endpoint:     conf.Endpoint,
		TLSEndpoint:  conf.TLSEndpoint,
		SMTPUsername: conf.SMTPUsername,
		SMTPPassword: conf.SMTPPassword,
		SMTPHost:     conf.SMTPHost,
		SMTPAddr:     conf.SMTPAddr,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.GetAllEventsHandler)
	s.Methods("DELETE").Path("").HandlerFunc(eh.DeleteAllEventsHandler)
	s.Methods("GET").Path("/bookings").HandlerFunc(eh.GetAllBookingsHandler)
	s.Methods("POST").Path("/bookings/{eventID}").HandlerFunc(eh.AddBookingHandler)
	s.Methods("DELETE").Path("/bookings/{bookingID}").HandlerFunc(eh.DeleteBookingHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	log.Printf("bookingservice listening to %s & %s...", eh.Endpoint, eh.TLSEndpoint)
	server := handlers.CORS()(r)
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", server)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, server)
	}()
	return httpErrChan, httpsErrChan
}
