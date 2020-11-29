package pickup

import (
	"fmt"
	"strings"

	"github.com/AlexCrane/tvslogparser/action"
)

type PickupAnalysis struct {
	byType map[action.PickupType]int
	total  int
}

func NewAnalysis(actions []action.Action) (*PickupAnalysis, error) {
	analysis := &PickupAnalysis{
		byType: make(map[action.PickupType]int),
		total:  0,
	}

	for _, a := range actions {
		if a.ActionType() == action.ActionTypePickup {
			pickup := a.(*action.Pickup)

			val, ok := analysis.byType[pickup.Type]
			if !ok {
				analysis.byType[pickup.Type] = 1
			} else {
				analysis.byType[pickup.Type] = val + 1
			}

			analysis.total++
		}
	}

	return analysis, nil
}

func (p *PickupAnalysis) FormatAsString() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Total pickups: %d\n", p.total))
	for _, k := range []action.PickupType{
		action.PTNone,
		action.PTFreeAttack,
		action.PTTripleAttack,
		action.PTExplosive50,
		action.PTExplosive250,
		action.PTExplosive450,
		action.PTRepair150,
		action.PTRepair200,
		action.PTRepair300,
		action.PTPoints100,
		action.PTPoints150,
		action.PTPoints250,
		action.PTPoints300,
		action.PTInvulnerability,
		action.PTHalfMaxRepair,
		action.PTGateIntel,
		action.PTDecoy,
		action.PTRandomHyperGate,
	} {
		v := p.byType[k]
		sb.WriteString(fmt.Sprintf("%-24s : %-4d (%03f%%)\n", k, v, float64(v)/float64(p.total)*100))
	}

	return sb.String()
}
