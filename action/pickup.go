package action

import (
	"fmt"
	"strconv"
	"strings"
)

type PickupType int

const (
	PTNone PickupType = iota
	PTFreeAttack
	PTTripleAttack
	PTExplosive50
	PTExplosive250
	PTExplosive450
	PTRepair150
	PTRepair200
	PTRepair300
	PTPoints100
	PTPoints150
	PTPoints250
	PTPoints300
	PTInvulnerability
	PTHalfMaxRepair
	PTGateIntel
	PTDecoy
	PTRandomHyperGate
)

var pickupTypeToString = map[PickupType]string{
	PTNone:            "None",
	PTFreeAttack:      "Free attack",
	PTTripleAttack:    "Triple attack",
	PTExplosive50:     "-50 hitpoints",
	PTExplosive250:    "-250 hitpoints",
	PTExplosive450:    "-450 hitpoints",
	PTRepair150:       "150 hitpoints",
	PTRepair200:       "200 hitpoints",
	PTRepair300:       "300 hitpoints",
	PTPoints100:       "100 points",
	PTPoints150:       "150 points",
	PTPoints250:       "250 points",
	PTPoints300:       "300 points",
	PTInvulnerability: "Invulnerability",
	PTHalfMaxRepair:   "Half max repair",
	PTGateIntel:       "Gate intel",
	PTDecoy:           "Decoy",
	PTRandomHyperGate: "Random hyper gate",
}

func (pt PickupType) String() string {
	s, ok := pickupTypeToString[pt]
	if ok {
		return s
	}
	return "<UNKNOWN>"
}

type Pickup struct {
	tick int
	Type PickupType
}

var _ Action = (*Pickup)(nil)

func (p *Pickup) String() string {
	return fmt.Sprintf("Pickup type %s on tick %d", p.Type, p.tick)
}

func (p *Pickup) ActionType() ActionType {
	return ActionTypePickup
}

func (p *Pickup) Tick() int {
	return p.tick
}

func PickupFromCSVRecord(record []string) (*Pickup, error) {
	tick, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, err
	}

	pickupText := record[3]

	// Strip of "Pickup ("
	innerText := pickupText[8:]
	// strip off trailing )
	innerText = innerText[:len(innerText)-1]

	for k, v := range pickupTypeToString {
		if innerText == v {
			return &Pickup{
				tick: tick,
				Type: k,
			}, nil
		}
	}

	if strings.HasPrefix(innerText, "Gate intel") {
		return &Pickup{
			tick: tick,
			Type: PTGateIntel,
		}, nil
	}

	return nil, fmt.Errorf("not implemented %q", innerText)
}
