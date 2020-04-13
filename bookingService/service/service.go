package service

import (
	"flag"
	"fmt"
	"github.com/agelloz/reach/bookingService/configuration"
	"github.com/agelloz/reach/bookingService/listener"
	"github.com/agelloz/reach/bookingService/persistence"
	"github.com/agelloz/reach/msgqueue"
	"github.com/agelloz/reach/msgqueue/msgqueue_amqp"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"net/http"
)

type BookingServiceHandler struct {
	DBHandler     persistence.DBHandler
	Endpoint      string
	TLSEndpoint   string
	EventListener msgqueue.EventListener
}

// ServeAPI is
func ServeAPI() (chan error, chan error) {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()
	conf, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to database...")
	dh, err := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(conf.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	el, err := msgqueue_amqp.NewAMQPEventListener(conn, "events_queue")
	if err != nil {
		panic(err)
	}
	processor := &listener.EventProcessor{EventListener: el, Database: dh}
	go processor.ProcessEvents()

	eh := &BookingServiceHandler{
		DBHandler:     dh,
		Endpoint:      conf.Endpoint,
		TLSEndpoint:   conf.TLSEndpoint,
		EventListener: el,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.GetAllEventsHandler)
	s.Methods("GET").Path("/bookings").HandlerFunc(eh.GetAllBookingsHandler)
	s.Methods("POST").Path("/bookings/{eventID}").HandlerFunc(eh.AddBookingHandler)
	s.Methods("DELETE").Path("/bookings/{eventID}").HandlerFunc(eh.DeleteBookingHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	fmt.Println("bookingService listening...")
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
