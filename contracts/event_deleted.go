package contracts

type EventDeletedEvent struct {
	ID [12]byte `json:"event_id"`
}

func (e *EventDeletedEvent) EventName() string {
	return "event.deleted"
}
