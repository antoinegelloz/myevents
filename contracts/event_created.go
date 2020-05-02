package contracts

import (
	"time"
)

type EventCreatedEvent struct {
	ID              [12]byte  `json:"event_id"`
	Name            string    `json:"event_name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	LocationName    string    `json:"location_name"`
	LocationCountry string    `json:"location_country"`
}

func (e *EventCreatedEvent) EventName() string {
	return "event.created"
}
