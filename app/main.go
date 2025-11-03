package main

import (
	"fmt"
	"time"

	"github.com/vladislav149/app/calendar"
	"github.com/vladislav149/app/events"
)

func main() {
	calendar.AddEvent("1", events.Event{
		Title:   "ass",
		StartAt: time.Now(),
	})

	fmt.Println("Работает")
}
