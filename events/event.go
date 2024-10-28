package models

import (
	"fmt"
	"time"
)

type Event struct {
	// gin supports struct tags
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func New() Event {
	// constructor
	return Event{}
}

func (ev Event) Print() {
	// method
	fmt.Println(ev.ID)

}
