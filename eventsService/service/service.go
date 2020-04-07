package service

import (
	"flag"
	"fmt"
	"github.com/agelloz/reach/eventsService/configuration"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/agelloz/reach/eventsService/persistence"
)

type EventsServiceHandler struct {
	dbHandler   persistence.DBHandler
	endpoint    string
	tlsEndpoint string
}

// ServeAPI is
func ServeAPI() (chan error, chan error) {
	confPath := flag.String("conf", `.\configuration\config.json`,
		"flag to set the path to the configuration json file")
	flag.Parse()
	conf, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database...")
	dh, _ := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	eh := &EventsServiceHandler{
		dbHandler:   dh,
		endpoint:    conf.Endpoint,
		tlsEndpoint: conf.TLSEndpoint,
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
		httpsErrChan <- http.ListenAndServeTLS(eh.tlsEndpoint, "cert.pem", "key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(eh.endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
