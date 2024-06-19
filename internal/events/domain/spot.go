package domain

import (
	"errors"
	"github.com/google/uuid"
)

type SpotStatus string

var(
	ErrInvalidSpotName = errors.New("invalid spot name")
	ErrSpotNotFound = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
	ErrSpotNameTwoCharacters = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameStartEndWithLetter = errors.New("spot name must start with a capital letter")
)

const (
	SpotStatusAvaliable SpotStatus = "avaliable"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func (s *Spot) validate() error {
	if len(s.Name) == 0 {
		return ErrEventNameRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartEndWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameStartEndWithLetter
	}


	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID: uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvaliable,
	}

	if err := spot.validate(); err != nil {
		return nil, err
	}

	return spot, nil
}