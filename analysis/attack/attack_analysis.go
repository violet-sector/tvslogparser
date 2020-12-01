package attack

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AlexCrane/tvslogparser/action"
)

const (
	bucketSize      = 50
	assumeMaxDamage = 5000
)

type Attack struct {
	sortedDamage       []int
	attackDistribution []int
	totalDamage        int
	totalAttacks       int
}

func NewAnalysis(actions []action.Action) (*Attack, error) {
	analysis := &Attack{
		attackDistribution: make([]int, assumeMaxDamage/bucketSize),
	}

	for _, a := range actions {
		if a.ActionType() == action.ActionTypeAttack {
			attack := a.(*action.Attack)
			if !attack.Outgoing {
				// TODO: Have separate structs for attacks/hits
				// for now, ignore hits on me
				continue
			}

			analysis.sortedDamage = append(analysis.sortedDamage, attack.Damage)
			analysis.attackDistribution[attack.Damage/bucketSize]++
			analysis.totalDamage += attack.Damage
			analysis.totalAttacks++
		}
	}

	sort.Ints(analysis.sortedDamage)

	maxDamage := analysis.sortedDamage[len(analysis.sortedDamage)-1]
	bucketTrim := (maxDamage / bucketSize) + 1
	analysis.attackDistribution = analysis.attackDistribution[:bucketTrim]

	return analysis, nil
}

func (p *Attack) FormatAsString() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Total attacks : %d\n", p.totalAttacks))
	sb.WriteString(fmt.Sprintf("Total damage  : %d\n", p.totalDamage))
	sb.WriteString(fmt.Sprintf("Mean damage   : %d\n", p.totalDamage/p.totalAttacks))
	sb.WriteString(fmt.Sprintf("Median damage : %d\n", p.sortedDamage[len(p.sortedDamage)/2]))
	sb.WriteString(fmt.Sprintf("Range         : %d->%d\n", p.sortedDamage[0], p.sortedDamage[len(p.sortedDamage)-1]))
	sb.WriteString("Attack Distribution:\n")
	for n := range p.attackDistribution {
		bucketRange := fmt.Sprintf("%d-%d", n*bucketSize, (n+1)*bucketSize)
		sb.WriteString(fmt.Sprintf("%-9s ", bucketRange))
	}
	sb.WriteRune('\n')
	for _, bucket := range p.attackDistribution {
		sb.WriteString(fmt.Sprintf("%-9d ", bucket))
	}
	sb.WriteRune('\n')

	return sb.String()
}
