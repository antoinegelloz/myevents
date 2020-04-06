package service

import (
	"net/http"

	"github.com/agelloz/reach/eventsService/persistence"

	"github.com/gorilla/mux"
)

type eventsServiceHandler struct {
	dbhandler persistence.DBHandler
}

// ServeAPI is
func ServeAPI(endpoint, tlsendpoint string, dh persistence.DBHandler) (chan error, chan error) {
	eh := &eventsServiceHandler{dbhandler: dh}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.getEventHandler)
	eventsrouter.Methods("DELETE").Path("/{nameOrID}/{nameOrIDValue}").HandlerFunc(eh.deleteEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(eh.getAllEventsHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(eh.addEventHandler)
	httpErrChan := make(chan error)
	httpsErrChan := make(chan error)
	go func() {
		httpsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r)
	}()
	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()
	return httpErrChan, httpsErrChan
}
