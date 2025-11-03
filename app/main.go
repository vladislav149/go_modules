package main

import (
	"app/calendar"
	"app/events"
	"time"
)

func main() {
	calendar.AddEvent("1", events.Event{
		Title:   "ass",
		StartAt: time.Now(),
	})
}
