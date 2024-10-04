package main

import (
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID   string
	Time time.Time
}

type DoorEvent struct {
	Event
	Action string //open, close
}

type TemperatureEvent struct {
	Event
	Value float32
}

func mainS() {

	evt, err := NewDoorEvent("front door", time.Now(), "open")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", evt)
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}
	evt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}
	return &evt, nil
}

/*

Type embedding is a technique where you include one type into another
but as a parameter without a name.

Through type embedding, all of the exported parameters and methods defined
on the embedded type are also accessible to the type within which the type
is embedded.






*/
