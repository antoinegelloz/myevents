package service

import (
	"flag"
	"fmt"
	"github.com/agelloz/reach/msgqueue"
	"github.com/streadway/amqp"
	"net/http"

	"github.com/agelloz/reach/eventsService/configuration"
	"github.com/agelloz/reach/eventsService/persistence"
	"github.com/agelloz/reach/msgqueue/msgqueue_amqp"
	"github.com/gorilla/mux"
)

type EventsServiceHandler struct {
	DbHandler    persistence.DBHandler
	Endpoint     string
	TLSEndpoint  string
	EventEmitter msgqueue.EventEmitter
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
	conn, err := amqp.Dial(conf.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	ee, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to database...")
	dh, err := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	if err != nil {
		panic(err)
	}
	eh := &EventsServiceHandler{
		DbHandler:    dh,
		Endpoint:     conf.Endpoint,
		TLSEndpoint:  conf.TLSEndpoint,
		EventEmitter: ee,
	}
	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.GetAllEventsHandler)
	s.Methods("POST").Path("").HandlerFunc(eh.AddEventHandler)
	s.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.GetEventHandler)
	s.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.DeleteEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	fmt.Println("eventsService listening...")
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
