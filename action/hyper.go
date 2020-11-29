package action

import (
	"fmt"
	"strconv"
	"strings"
)

type Hyper struct {
	tick      int
	Location  string
	Autopilot bool
}

var _ Action = (*Hyper)(nil)

func (h *Hyper) String() string {
	if h.Autopilot {
		return fmt.Sprintf("Hypered to %s (autopilot) on tick %d", h.Location, h.tick)
	}
	return fmt.Sprintf("Hypered to %s on tick %d", h.Location, h.tick)
}

func (h *Hyper) ActionType() ActionType {
	return ActionTypeHyper
}

func (h *Hyper) Tick() int {
	return h.tick
}

func HyperFromCSVRecord(record []string) (*Hyper, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	return &Hyper{
		tick:      tick,
		Location:  strings.TrimSuffix(strings.TrimPrefix(record[3], "Hypered to "), " (via auto-pilot)"),
		Autopilot: strings.HasSuffix(record[3], " (via auto-pilot)"),
	}, nil
}
