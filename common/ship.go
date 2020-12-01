package common

import "fmt"

const (
	CruiserShipType = "cruiser/carrier"
)

type ShipLevel struct {
	Level             int
	HasCombatPiloting bool
	CombatDelta       int
	PilotingDelta     int
}

func (s *ShipLevel) String() string {
	if s.IsCruiser() {
		return fmt.Sprintf("cruiser")
	}
	if s.HasCombatPiloting {
		return fmt.Sprintf("level %d (adding %d combat and %d piloting)", s.Level, s.CombatDelta, s.PilotingDelta)
	}
	return fmt.Sprintf("level %d", s.Level)
}

func (s *ShipLevel) IsCruiser() bool {
	return s.Level == 6
}

type ShipStats struct {
	shipType               string
	level                  int
	combatPilotingReliable bool
	combat                 int
	piloting               int
}

func (s *ShipStats) LevelUp(l ShipLevel) {
	s.level = l.Level
	if l.HasCombatPiloting {
		s.piloting += l.PilotingDelta
		s.combat += l.CombatDelta
	} else {
		s.combatPilotingReliable = false
	}

	if s.level == 6 {
		s.shipType = CruiserShipType
	}
}
