package templater

import (
	"errors"
)

type EventChannel struct {
	Channel chan Event
}

func (ec *EventChannel) Send(event Event) (Event, error) {
	// HERE
	return event, nil
}

var (
	channelCreated bool = false
	channel        chan Event
)

func CreateChannel() chan Event {
	if !channelCreated {
		channelCreated = true
		channel = make(chan Event)
	}

	return channel
}

func SendEvent(event Event) (Event, error) {
	if !channelCreated {
		return event, errors.New("channel is not initialized")
	}

	channel <- event
	return event, nil
}
