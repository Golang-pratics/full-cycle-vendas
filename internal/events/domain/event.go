package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateRequired = errors.New("event date is required")
	ErrEventCapacityRequired = errors.New("event capacity is required")
	ErrEventPriceRequired = errors.New("event price is required")

)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)



type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartinerID   int
	Spots        []Spot
	Tickets      []Ticket
}

func (e *Event) validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateRequired
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityRequired
	}

	if e.Price <= 0 {
		return ErrEventPriceRequired
	}

	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)

	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)

	return spot, nil
}

func (s *Spot) ReserveSpot(ticket *Ticket) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusSold
	s.TicketID = ticket.ID

	return nil
}