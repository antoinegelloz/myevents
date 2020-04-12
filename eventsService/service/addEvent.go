package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/agelloz/reach/eventsService/contracts"
	"net/http"
	"time"

	"github.com/agelloz/reach/eventsService/models"
)

func (eh *EventsServiceHandler) addEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Cannot decode event data", http.StatusInternalServerError)
		return
	}
	id, err := eh.dbHandler.AddEvent(event)
	if nil != err {
		http.Error(w, fmt.Sprintf("Cannot add event ID: %s", id), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Added new event ID:%d\n", id)

	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: event.Location.ID.String(),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}
	err = eh.eventEmitter.Emit(&msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot emit event ID: %s", id), http.StatusInternalServerError)
		return
	}
	fmt.Print("New event successfully emitted\n")
}
