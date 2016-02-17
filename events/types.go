package events

import "time"

//Event represents an event sent to all Observers as part of the Observer pattern
type Event interface {
	//When the event occurred
	When() time.Time
}

//Observer represents the registerable Observer as part of the Observer pattern
type Observer interface {
	Receive(e Event)
}

//Observable represents any observable type as per the Observer pattern
type Observable interface {
	Add(o Observer)
	Remove(o Observer)
	Notify() error
}

type appEvent struct {
	when time.Time
}

func (e appEvent) When() time.Time {
	return e.when
}
