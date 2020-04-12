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
	dbHandler    persistence.DBHandler
	endpoint     string
	tlsEndpoint  string
	eventEmitter msgqueue.EventEmitter
}

// ServeAPI is
func ServeAPI() (chan error, chan error) {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()
	conf, _ := configuration.ExtractConfiguration(*confPath)

	conn, err := amqp.Dial(conf.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	ee, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to database...")
	dh, _ := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	eh := &EventsServiceHandler{
		dbHandler:    dh,
		endpoint:     conf.Endpoint,
		tlsEndpoint:  conf.TLSEndpoint,
		eventEmitter: ee,
	}
	r := mux.NewRouter()
	s := r.PathPrefix("/events").Subrouter()
	s.Methods("GET").Path("").HandlerFunc(eh.getAllEventsHandler)
	s.Methods("POST").Path("").HandlerFunc(eh.addEventHandler)
	s.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.getEventHandler)
	s.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.deleteEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	fmt.Println("eventsService listening...")
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.tlsEndpoint, "certificate/cert.pem", "certificate/key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
