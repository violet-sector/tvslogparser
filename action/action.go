package action

import (
	"errors"
	"fmt"
	"strings"
)

type ActionType int

const (
	ActionTypeUnknown ActionType = iota
	ActionTypePickup
)

type Action interface {
	fmt.Stringer
	ActionType() ActionType
	Tick() int
}

func NewActionFromCSVRecord(record []string) (Action, error) {
	if len(record) != 6 {
		return nil, fmt.Errorf("unexpected record len %d", len(record))
	}

	if strings.HasPrefix(record[3], "Pickup") {
		return PickupFromCSVRecord(record)
	}

	return nil, errors.New("not implemented")
}
