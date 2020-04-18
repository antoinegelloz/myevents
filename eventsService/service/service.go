package service

import (
	"flag"
	"github.com/streadway/amqp"
	"log"
	"net/http"

	"github.com/agelloz/reach/eventsService/configuration"
	"github.com/agelloz/reach/eventsService/persistence"
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
	conn, err := amqp.Dial(conf.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	_, err = channel.QueueDeclare("events_queue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	log.Println("connecting to database...")
	dh, err := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	if err != nil {
		panic(err)
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
	s.Methods("POST").Path("").HandlerFunc(eh.AddEventHandler)
	s.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.GetEventHandler)
	s.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.DeleteEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	log.Println("eventsService listening...")
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.TLSEndpoint, "certificate/cert.pem", "certificate/key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.Endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
