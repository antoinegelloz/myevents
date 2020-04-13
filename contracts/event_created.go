package contracts

import (
	"time"
)

type EventCreatedEvent struct {
	ID         []byte    `json:"event_id"`
	Name       string    `json:"event_name"`
	LocationID []byte    `json:"location_id"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

func (e *EventCreatedEvent) EventName() string {
	return "event.created"
}
