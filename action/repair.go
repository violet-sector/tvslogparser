package action

import (
	"fmt"
	"strconv"

	"github.com/AlexCrane/tvslogparser/common"
)

type Repair struct {
	tick               int
	Outgoing           bool
	FriendlyPilot      string
	Hitpoints          int
	HitpointsRemaining int
}

var _ Action = (*Repair)(nil)

func (r *Repair) String() string {
	if r.Outgoing {
		return fmt.Sprintf("Repaired %s by %d hitpoints (%d remaining) on tick %d", r.FriendlyPilot, r.Hitpoints, r.HitpointsRemaining, r.tick)
	}
	return fmt.Sprintf("%s repaired me by %d hitpoints (%d remaining) on tick %d", r.FriendlyPilot, r.Hitpoints, r.HitpointsRemaining, r.tick)
}

func (r *Repair) ActionType() ActionType {
	return ActionTypeRepair
}

func (r *Repair) Tick() int {
	return r.tick
}

func RepairFromCSVRecord(record []string, myPlayer *common.Player) (*Repair, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	var giver, receiver string
	matches, err := fmt.Sscanf(record[3], "%s repaired %s", &giver, &receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to parse repair action: %w", err)
	}
	if matches != 2 {
		return nil, fmt.Errorf("unexpected matches in repair action %d", matches)
	}

	hp, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse repair hp: %w", err)
	}

	remaining, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse repair hitpoints remaining: %w", err)
	}

	if giver == myPlayer.GetName() {
		return &Repair{
			tick:               tick,
			Outgoing:           true,
			FriendlyPilot:      receiver,
			Hitpoints:          hp,
			HitpointsRemaining: remaining,
		}, nil
	}

	return &Repair{
		tick:               tick,
		Outgoing:           false,
		FriendlyPilot:      giver,
		Hitpoints:          hp,
		HitpointsRemaining: remaining,
	}, nil
}
