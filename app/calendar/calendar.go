package calendar

import "github.com/vladislav149/app/events"

var calendar = make(map[string]events.Event)

func AddEvent(name string, event events.Event) {
	calendar[name] = event
}
