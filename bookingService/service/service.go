package service

import (
	"flag"
	"log"
	"net/http"

	"github.com/agelloz/reach/bookingService/configuration"
	"github.com/agelloz/reach/bookingService/listener"
	"github.com/agelloz/reach/bookingService/persistence"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type BookingServiceHandler struct {
	DBHandler   persistence.DBHandler
	Endpoint    string
	TLSEndpoint string
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
	log.Println("connecting to database...")
	dh, err := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	if err != nil {
		log.Panic(err)
	}
	go listener.Listen(conf.AMQPMessageBroker, dh)

	eh := &BookingServiceHandler{
		DBHandler:   dh,
		Endpoint:    conf.Endpoint,
		TLSEndpoint: conf.TLSEndpoint,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.GetAllEventsHandler)
	s.Methods("DELETE").Path("").HandlerFunc(eh.DeleteAllEventsHandler)
	s.Methods("GET").Path("/bookings").HandlerFunc(eh.GetAllBookingsHandler)
	s.Methods("POST").Path("/bookings/{eventID}").HandlerFunc(eh.AddBookingHandler)
	s.Methods("DELETE").Path("/bookings/{eventID}").HandlerFunc(eh.DeleteBookingHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	log.Println("bookingService listening...")
	server := handlers.CORS()(r)
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", server)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, server)
	}()
	return httpErrChan, httpsErrChan
}
