package action

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AlexCrane/tvslogparser/common"
)

type Repair struct {
	tick                     int
	Outgoing                 bool
	FriendlyPilot            string
	TargetHitpoints          int
	TargetHitpointsRemaining int

	// Round 43+
	PlayerDamageTaken        int
	PlayerHitpointsRemaining int
}

var _ Action = (*Repair)(nil)

func (r *Repair) String() string {
	if r.Outgoing {
		return fmt.Sprintf("Repaired %s by %d hitpoints (%d remaining) on tick %d", r.FriendlyPilot, r.TargetHitpoints, r.TargetHitpointsRemaining, r.tick)
	}
	return fmt.Sprintf("%s repaired me by %d hitpoints (%d remaining) on tick %d", r.FriendlyPilot, r.TargetHitpoints, r.TargetHitpointsRemaining, r.tick)
}

func (r *Repair) ActionType() ActionType {
	return ActionTypeRepair
}

func (r *Repair) Tick() int {
	return r.tick
}

func (r *Repair) IsSuicidalRepair() bool {
	return r.PlayerHitpointsRemaining == -1
}

var repairRegex = regexp.MustCompile("(.*) repaired (.*)")

func RepairFromCSVRecord(record []string, myPlayer *common.Player) (*Repair, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	matches := repairRegex.FindStringSubmatch(record[3])
	if matches == nil {
		return nil, fmt.Errorf("failed to parse repair action")
	}
	if len(matches) != 3 {
		return nil, fmt.Errorf("unexpected matches in repair action %d: %v", len(matches), matches)
	}
	giver := matches[1]
	receiver := matches[2]

	hp, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse repair hp: %w", err)
	}

	remaining, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse repair hitpoints remaining: %w", err)
	}

	var playerDamage int
	var playerRemaining int
	if len(record) >= 8 && giver == myPlayer.GetName() {
		var err error
		// This is a round 43+ record - it has two additional fields - "Player damage taken","Player HP remaining"
		playerDamage, err = strconv.Atoi(record[6])
		if err != nil {
			return nil, fmt.Errorf("failed to parse attack player damage: %w", err)
		}
		if playerDamage < 0 {
			playerDamage = -playerDamage
		}

		playerRemaining, err = strconv.Atoi(record[7])
		if err != nil {
			return nil, fmt.Errorf("failed to parse attack player hitpoints remaining: %w", err)
		}
	}

	if giver == myPlayer.GetName() {
		return &Repair{
			tick:                     tick,
			Outgoing:                 true,
			FriendlyPilot:            receiver,
			TargetHitpoints:          hp,
			TargetHitpointsRemaining: remaining,
			PlayerDamageTaken:        playerDamage,
			PlayerHitpointsRemaining: playerRemaining,
		}, nil
	}

	return &Repair{
		tick:                     tick,
		Outgoing:                 false,
		FriendlyPilot:            giver,
		TargetHitpoints:          hp,
		TargetHitpointsRemaining: remaining,
		PlayerDamageTaken:        playerDamage,
		PlayerHitpointsRemaining: playerRemaining,
	}, nil
}
