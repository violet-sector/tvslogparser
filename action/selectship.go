package action

import (
	"fmt"
	"strconv"
	"strings"
)

type SelectShip struct {
	tick     int
	ShipType string
}

var _ Action = (*SelectShip)(nil)

func (s *SelectShip) String() string {
	return fmt.Sprintf("Selected ship %s on tick %d", s.ShipType, s.tick)
}

func (s *SelectShip) ActionType() ActionType {
	return ActionTypeSelectShip
}

func (s *SelectShip) Tick() int {
	return s.tick
}

func SelectShipFromCSVRecord(record []string) (*SelectShip, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	return &SelectShip{
		tick:     tick,
		ShipType: strings.TrimPrefix(record[3], "Selected to pilot a "),
	}, nil
}
