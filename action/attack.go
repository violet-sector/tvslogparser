package action

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/AlexCrane/tvslogparser/common"
)

type Attack struct {
	tick                     int
	Outgoing                 bool
	EnemyPilot               string
	TargetDamage             int
	TargetHitpointsRemaining int

	// Round 43+
	PlayerDamageTaken        int
	PlayerHitpointsRemaining int
}

var _ Action = (*Attack)(nil)

func (a *Attack) String() string {
	if a.Outgoing {
		return fmt.Sprintf("Atacked %s for %d damage leaving %d remaining on tick %d", a.EnemyPilot, a.TargetDamage, a.TargetHitpointsRemaining, a.tick)
	}
	return fmt.Sprintf("Atacked by %s for %d damage leaving %d remaining on tick %d", a.EnemyPilot, a.TargetDamage, a.TargetHitpointsRemaining, a.tick)
}

func (a *Attack) ActionType() ActionType {
	return ActionTypeAttack
}

func (a *Attack) Tick() int {
	return a.tick
}

func (a *Attack) IsKill() bool {
	return a.Outgoing && a.TargetHitpointsRemaining == -1
}

func (a *Attack) IsDeath() bool {
	return !a.Outgoing && a.TargetHitpointsRemaining == -1
}

func (a *Attack) IsSuicidalAttack() bool {
	return a.PlayerHitpointsRemaining == -1
}

var attackRegex = regexp.MustCompile("(.*) attacked (.*)")

func AttackFromCSVRecord(record []string, myPlayer *common.Player) (*Attack, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	matches := attackRegex.FindStringSubmatch(record[3])
	if matches == nil {
		return nil, fmt.Errorf("failed to parse attack action")
	}
	if len(matches) != 3 {
		return nil, fmt.Errorf("unexpected matches in attack action %d: %v", len(matches), matches)
	}
	attacker := matches[1]
	defender := matches[2]

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

	var playerDamage int
	var playerRemaining int
	if len(record) >= 8 && attacker == myPlayer.GetName() {
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

	if attacker == myPlayer.GetName() {
		return &Attack{
			tick:                     tick,
			Outgoing:                 true,
			EnemyPilot:               defender,
			TargetDamage:             damage,
			TargetHitpointsRemaining: remaining,
			PlayerDamageTaken:        playerDamage,
			PlayerHitpointsRemaining: playerRemaining,
		}, nil
	}

	return &Attack{
		tick:                     tick,
		Outgoing:                 false,
		EnemyPilot:               attacker,
		TargetDamage:             damage,
		TargetHitpointsRemaining: remaining,
		PlayerDamageTaken:        playerDamage,
		PlayerHitpointsRemaining: playerRemaining,
	}, nil
}
