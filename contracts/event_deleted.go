package contracts

type EventDeletedEvent struct {
	ID         string    `json:"id"`
}

func (e *EventDeletedEvent) EventName() string {
	return "event.deleted"
}
