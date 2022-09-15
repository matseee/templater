package communication

import "errors"

func CreateEventChannel() EventChannel {
	return EventChannel{
		Channel: make(chan Event),
	}
}

type EventChannel struct {
	Channel chan Event
}

func (ec *EventChannel) Send(event Event) error {
	if err := ec.checkChannel(); err != nil {
		return err
	}

	ec.Channel <- event
	return nil
}

func (ec *EventChannel) checkChannel() error {
	if ec.Channel == nil {
		return errors.New("channel is closed or nil")
	}
	return nil
}
