package service

import (
	"net/http"
)

func (eh *BookingServiceHandler) DeleteAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	eh.DBHandler.DeleteAllEvents()
}
