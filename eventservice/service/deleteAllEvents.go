package service

import (
	"net/http"
)

func (eh *EventsServiceHandler) DeleteAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	eh.DbHandler.DeleteAllEvents()
}
