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

func (pt PickupType) String() string {
	switch pt {
	case PTNone:
		return "None"
	case PTFreeAttack:
		return "Free attack"
	case PTTripleAttack:
		return "Triple attack"
	case PTExplosive50:
		return "-50 hitpoints"
	case PTExplosive250:
		return "-250 hitpoints"
	case PTExplosive450:
		return "-450 hitpoints"
	case PTRepair150:
		return "150 hitpoints"
	case PTRepair200:
		return "200 hitpoints"
	case PTRepair300:
		return "300 hitpoints"
	case PTPoints100:
		return "100 points"
	case PTPoints150:
		return "150 points"
	case PTPoints250:
		return "250 points"
	case PTPoints300:
		return "300 points"
	case PTInvulnerability:
		return "Invulnerability"
	case PTHalfMaxRepair:
		return "Half max repair"
	case PTGateIntel:
		return "Gate intel"
	case PTDecoy:
		return "Decoy"
	case PTRandomHyperGate:
		return "Random hyper gate"
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

	var ptype PickupType
	switch innerText {
	case "None":
		ptype = PTNone
	case "Free attack":
		ptype = PTFreeAttack
	case "Triple attack":
		ptype = PTTripleAttack
	case "-50 hitpoints":
		ptype = PTExplosive50
	case "-250 hitpoints":
		ptype = PTExplosive250
	case "-450 hitpoints":
		ptype = PTExplosive450
	case "150 hitpoints":
		ptype = PTRepair150
	case "200 hitpoints":
		ptype = PTRepair200
	case "300 hitpoints":
		ptype = PTRepair300
	case "100 points":
		ptype = PTPoints100
	case "150 points":
		ptype = PTPoints150
	case "250 points":
		ptype = PTPoints250
	case "300 points":
		ptype = PTPoints300
	case "Invulnerability":
		ptype = PTInvulnerability
	case "Half max repair":
		ptype = PTHalfMaxRepair
	case "Random hyper gate":
		ptype = PTRandomHyperGate
	case "Decoy":
		ptype = PTDecoy
	default:
		if strings.HasPrefix(innerText, "Gate intel") {
			ptype = PTGateIntel
			break
		}

		return nil, fmt.Errorf("not implemented %q", innerText)
	}

	return &Pickup{
		tick: tick,
		Type: ptype,
	}, nil
}
