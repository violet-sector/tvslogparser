package action

import (
	"fmt"
	"strconv"

	"github.com/AlexCrane/tvslogparser/common"
)

type Attack struct {
	tick               int
	Outgoing           bool
	EnemyPilot         string
	Damage             int
	HitpointsRemaining int
}

var _ Action = (*Attack)(nil)

func (a *Attack) String() string {
	if a.Outgoing {
		return fmt.Sprintf("Atacked %s for %d damage leaving %d remaining on tick %d", a.EnemyPilot, a.Damage, a.HitpointsRemaining, a.tick)
	}
	return fmt.Sprintf("Atacked by %s for %d damage leaving %d remaining on tick %d", a.EnemyPilot, a.Damage, a.HitpointsRemaining, a.tick)
}

func (a *Attack) ActionType() ActionType {
	return ActionTypeAttack
}

func (a *Attack) Tick() int {
	return a.tick
}

func (a *Attack) IsKill() bool {
	return a.Outgoing && a.HitpointsRemaining == -1
}

func (a *Attack) IsDeath() bool {
	return !a.Outgoing && a.HitpointsRemaining == -1
}

func AttackFromCSVRecord(record []string, myPlayer *common.Player) (*Attack, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	var attacker, defender string
	matches, err := fmt.Sscanf(record[3], "%s attacked %s", &attacker, &defender)
	if err != nil {
		return nil, fmt.Errorf("failed to parse attack action: %w", err)
	}
	if matches != 2 {
		return nil, fmt.Errorf("unexpected matches in attack action %d", matches)
	}

	damage, err := strconv.Atoi(record[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse attack damage: %w", err)
	}
	if damage < 0 {
		damage = -damage
	}

	remaining, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, fmt.Errorf("failed to parse attack hitpoints remaining: %w", err)
	}

	if attacker == myPlayer.GetName() {
		return &Attack{
			tick:               tick,
			Outgoing:           true,
			EnemyPilot:         defender,
			Damage:             damage,
			HitpointsRemaining: remaining,
		}, nil
	}

	return &Attack{
		tick:               tick,
		Outgoing:           false,
		EnemyPilot:         attacker,
		Damage:             damage,
		HitpointsRemaining: remaining,
	}, nil
}
