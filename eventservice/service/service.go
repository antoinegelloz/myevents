package service

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"

	"github.com/agelloz/myevents/eventservice/configuration"
	"github.com/agelloz/myevents/eventservice/persistence"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type EventsServiceHandler struct {
	DbHandler      persistence.DBHandler
	Endpoint       string
	TLSEndpoint    string
	AMQPConnection *amqp.Connection
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

	var conn *amqp.Connection
	log.Println("connecting to AMQP message broker...")
	conn, err = amqp.Dial(conf.AMQPMessageBroker)
	for err != nil {
		log.Printf("AMQP connection error: %s\n", err)
		time.Sleep(2000000000)
		conn, err = amqp.Dial(conf.AMQPMessageBroker)
	}

	var dh persistence.DBHandler
	log.Println("connecting to database...")
	dh, err = persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	for err != nil {
		log.Printf("database connection error: %s\n", err)
		time.Sleep(2000000000)
		dh, err = persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	}

	eh := &EventsServiceHandler{
		DbHandler:      dh,
		Endpoint:       conf.Endpoint,
		TLSEndpoint:    conf.TLSEndpoint,
		AMQPConnection: conn,
	}
	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.GetAllEventsHandler)
	s.Methods("DELETE").Path("").HandlerFunc(eh.DeleteAllEventsHandler)
	s.Methods("POST").Path("").HandlerFunc(eh.AddEventHandler)
	s.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.GetEventHandler)
	s.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.DeleteEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	log.Println("eventsService listening...")
	server := handlers.CORS()(r)
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", server)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, server)
	}()
	return httpErrChan, httpsErrChan
}
