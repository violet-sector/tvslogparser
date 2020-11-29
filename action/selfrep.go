package action

import (
	"fmt"
	"strconv"
)

type SelfRep struct {
	tick int
}

var _ Action = (*SelfRep)(nil)

func (s *SelfRep) String() string {
	return fmt.Sprintf("Self rep on tick %d", s.tick)
}

func (s *SelfRep) ActionType() ActionType {
	return ActionTypeSelfRep
}

func (s *SelfRep) Tick() int {
	return s.tick
}

func SelfRepFromCSVRecord(record []string) (*SelfRep, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	return &SelfRep{
		tick: tick,
	}, nil
}
