package contracts

type EventDeletedEvent struct {
	ID []byte `json:"event_id"`
}

func (e *EventDeletedEvent) EventName() string {
	return "event.deleted"
}
