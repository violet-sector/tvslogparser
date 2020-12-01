package common

const (
	CruiserShipType = "cruiser/carrier"
)

type ShipLevel struct {
	Level             int
	HasCombatPiloting bool
	CombatDelta       int
	PilotingDelta     int
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

func (s *ShipStats) IsCruiser() bool {
	return s.level == 6
}
