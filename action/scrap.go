package action

import (
	"fmt"
	"strconv"
)

type ScrapActionType int

const (
	ScrapActionTypeUnknown ScrapActionType = iota
	ScrapActionTypeRetrieved
	ScrapActionTypeDumped
)

type Scrap struct {
	tick   int
	Type   ScrapActionType
	Amount int
}

var _ Action = (*Scrap)(nil)

func (s *Scrap) String() string {
	switch s.Type {
	case ScrapActionTypeRetrieved:
		return fmt.Sprintf("Retrieved %d scrap on tick %d", s.Amount, s.tick)
	case ScrapActionTypeDumped:
		return fmt.Sprintf("Dumped %d scrap on tick %d", s.Amount, s.tick)
	default:
		return "<unknown scrap action>"
	}
}

func (s *Scrap) ActionType() ActionType {
	return ActionTypeScrap
}

func (s *Scrap) Tick() int {
	return s.tick
}

func typeFromTypeString(s string) ScrapActionType {
	switch s {
	case "Retrieved":
		return ScrapActionTypeRetrieved
	case "Dumped":
		return ScrapActionTypeDumped
	default:
		return ScrapActionTypeUnknown
	}
}

func ScrapFromCSVRecord(record []string) (*Scrap, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	var stype string
	var amount int
	matches, err := fmt.Sscanf(record[3], "%s %d scrap", &stype, &amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse scrap action: %w", err)
	}
	if matches != 2 {
		return nil, fmt.Errorf("unexpected matches in scrap action %d", matches)
	}

	return &Scrap{
		tick: tick,
		Type: typeFromTypeString(stype),
	}, nil
}
