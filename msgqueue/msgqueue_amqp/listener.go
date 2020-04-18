package msgqueue_amqp

import (
	"encoding/json"
	"fmt"
	"github.com/agelloz/reach/contracts"
	"github.com/agelloz/reach/msgqueue"
	"github.com/streadway/amqp"
	"log"
)

type amqpEventListener struct {
	connection *amqp.Connection
	queue      string
}

func NewAMQPEventListener(conn *amqp.Connection, queue string) (msgqueue.EventListener, error) {
	listener := &amqpEventListener{
		connection: conn,
		queue:      queue,
	}
	err := listener.setup()
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func (a *amqpEventListener) setup() error {
	channel, err := a.connection.Channel()
	if err != nil {
		return nil
	}
	defer channel.Close()
	_, err = channel.QueueDeclare(a.queue, true, false, false, false, nil)
	return err
}

func (a *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	channel, err := a.connection.Channel()
	if err != nil {
		log.Println("channel error")
		return nil, nil, err
	}
	defer channel.Close()
	for _, eventName := range eventNames {
		if err := channel.QueueBind(a.queue, eventName, "eventsExchange", false, nil); err != nil {
			log.Println("queue binding error")
			return nil, nil, err
		}
	}
	messages, err := channel.Consume(a.queue, "", false, false, false, false, nil)
	if err != nil {
		log.Println("channel consuming error")
		return nil, nil, err
	}
	eventsChan := make(chan msgqueue.Event)
	errorsChan := make(chan error)
	log.Println("listening to events...")
	go func() {
		for msg := range messages {
			log.Println("message received")
			// Map message to actual event struct
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errorsChan <- fmt.Errorf("msg did not contain x-event-name header")
				err := msg.Nack(false, false)
				if err != nil {
					errorsChan <- err
				}
				continue
			}
			eventName, ok := rawEventName.(string)
			if !ok {
				errorsChan <- fmt.Errorf("x-event-name header is not string, but %t", rawEventName)
				err := msg.Nack(false, false)
				if err != nil {
					errorsChan <- err
				}
				continue
			}
			var event msgqueue.Event
			switch eventName {
			case "event.created":
				event = new(contracts.EventCreatedEvent)
			case "event.deleted":
				event = new(contracts.EventDeletedEvent)
			default:
				errorsChan <- fmt.Errorf("event type %s is unknown", eventName)
				continue
			}
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errorsChan <- err
				continue
			}
			log.Println("event sent to channel")
			eventsChan <- event
		}
	}()
	return eventsChan, errorsChan, nil
}
