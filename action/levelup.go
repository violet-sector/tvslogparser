package action

import (
	"fmt"
	"strconv"
	"strings"
)

type LevelUp struct {
	tick              int
	Level             int
	HasCombatPiloting bool
	CombatDelta       int
	PilotingDelta     int
}

var _ Action = (*LevelUp)(nil)

func (l *LevelUp) String() string {
	if l.Level == 6 {
		return fmt.Sprintf("Acheived cruiser on tick %d", l.tick)
	}
	return fmt.Sprintf("Acheived level %d on tick %d", l.Level, l.tick)
}

func (l *LevelUp) ActionType() ActionType {
	return ActionTypeLevelUp
}

func (l *LevelUp) Tick() int {
	return l.tick
}

func (l *LevelUp) IsCruiser() bool {
	return l.Level == 6
}

func LevelUpFromCSVRecord(record []string) (*LevelUp, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	if strings.Contains(record[3], "cruiser") {
		return &LevelUp{
			tick:  tick,
			Level: 6,
		}, nil
	}

	var level, combat, piloting int
	matches, err := fmt.Sscanf(record[3], "Achieved level %d (added %d combat / %d piloting)", &level, &combat, &piloting)
	if err != nil {
		matches, err = fmt.Sscanf(record[3], "Achieved level %d", &level)
		if err != nil {
			return nil, fmt.Errorf("failed to parse level up: %w", err)
		}
		if matches != 1 {
			return nil, fmt.Errorf("unexpected matches in level up %d", matches)
		}

		return &LevelUp{
			tick:              tick,
			Level:             level,
			HasCombatPiloting: false,
		}, nil
	}
	if matches != 3 {
		return nil, fmt.Errorf("unexpected matches in level up %d", matches)
	}

	return &LevelUp{
		tick:              tick,
		Level:             level,
		HasCombatPiloting: true,
		CombatDelta:       combat,
		PilotingDelta:     piloting,
	}, nil
}
