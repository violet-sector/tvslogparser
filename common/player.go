package common

type Player struct {
	name string
	ship ShipStats
}

func NewPlayer(pilotName string) *Player {
	return &Player{
		name: pilotName,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetShip() *ShipStats {
	return &p.ship
}

func (p *Player) SelectShip(shipType string) {
	p.ship = ShipStats{
		shipType:               shipType,
		level:                  1,
		combatPilotingReliable: true,
		combat:                 10,
		piloting:               10,
	}
}

func (p *Player) SetShip(shipType string) {
	p.ship.shipType = shipType
}
