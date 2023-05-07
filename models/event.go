package models

import "time"

type EventType int

const (
	OrderBookUpdate   EventType = iota
	Trade             EventType = iota
	OrderBookSnapshot EventType = iota
)

type Event struct {
	Timestamp time.Time
	Type      EventType
	Data      string
}

type Order struct {
	Price  float64
	Volume float64
}
