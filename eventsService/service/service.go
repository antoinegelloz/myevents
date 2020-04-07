package service

import (
	"flag"
	"fmt"
	"github.com/agelloz/reach/eventsService/configuration"
	"net/http"

	"github.com/agelloz/reach/eventsService/persistence"
	"github.com/gorilla/mux"
)

type EventsServiceHandler struct {
	dbHandler   persistence.DBHandler
	endpoint    string
	tlsEndpoint string
}

// ServeAPI is
func ServeAPI() (chan error, chan error) {
	eh := NewEventsServiceHandler()
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.getEventHandler)
	eventsRouter.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.deleteEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(eh.getAllEventsHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(eh.addEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	fmt.Println("eventsService listening...")
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(eh.tlsEndpoint, "cert.pem", "key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}

// NewEventsServiceHandler is
func NewEventsServiceHandler() *EventsServiceHandler {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()
	conf, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database...")
	dh, _ := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	return &EventsServiceHandler{
		dbHandler:   dh,
		endpoint:    conf.Endpoint,
		tlsEndpoint: conf.TLSEndpoint,
	}
}
