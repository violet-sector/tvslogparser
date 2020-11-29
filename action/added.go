package action

import (
	"fmt"
	"strconv"
)

type Added struct {
	tick      int
	PilotName string
}

var _ Action = (*Added)(nil)

func (a *Added) String() string {
	return fmt.Sprintf("Started game as pilot %s on tick %d", a.PilotName, a.tick)
}

func (a *Added) ActionType() ActionType {
	return ActionTypeAdded
}

func (a *Added) Tick() int {
	return a.tick
}

func AddedFromCSVRecord(record []string) (*Added, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	return &Added{
		tick:      tick,
		PilotName: record[1],
	}, nil
}
