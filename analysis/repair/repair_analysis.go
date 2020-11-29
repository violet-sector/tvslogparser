package repair

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AlexCrane/tvslogparser/action"
)

const (
	bucketSize      = 50
	assumeMaxRepair = 2000
)

type Repair struct {
	sortedRepHP     []int
	repDistribution []int
	totalRepHP      int
	totalReps       int
}

func NewAnalysis(actions []action.Action) (*Repair, error) {
	analysis := &Repair{
		repDistribution: make([]int, assumeMaxRepair/bucketSize),
	}

	for _, a := range actions {
		if a.ActionType() == action.ActionTypeRepair {
			repair := a.(*action.Repair)

			analysis.sortedRepHP = append(analysis.sortedRepHP, repair.Hitpoints)
			analysis.repDistribution[repair.Hitpoints/bucketSize]++
			analysis.totalRepHP += repair.Hitpoints
			analysis.totalReps++
		}
	}

	sort.Ints(analysis.sortedRepHP)

	maxDamage := analysis.sortedRepHP[len(analysis.sortedRepHP)-1]
	bucketTrim := (maxDamage / bucketSize) + 1
	analysis.repDistribution = analysis.repDistribution[:bucketTrim]

	return analysis, nil
}

func (p *Repair) FormatAsString() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Total repairs : %d\n", p.totalReps))
	sb.WriteString(fmt.Sprintf("Total repped  : %d\n", p.totalRepHP))
	sb.WriteString(fmt.Sprintf("Mean repair   : %d\n", p.totalRepHP/p.totalReps))
	sb.WriteString(fmt.Sprintf("Median repair : %d\n", p.sortedRepHP[len(p.sortedRepHP)/2]))
	sb.WriteString(fmt.Sprintf("Range         : %d->%d\n", p.sortedRepHP[0], p.sortedRepHP[len(p.sortedRepHP)-1]))
	sb.WriteString("Repair Distribution:\n")
	for n, _ := range p.repDistribution {
		bucketRange := fmt.Sprintf("%d-%d", n*bucketSize, (n+1)*bucketSize)
		sb.WriteString(fmt.Sprintf("%-9s ", bucketRange))
	}
	sb.WriteRune('\n')
	for _, bucket := range p.repDistribution {
		sb.WriteString(fmt.Sprintf("%-9d ", bucket))
	}
	sb.WriteRune('\n')

	return sb.String()
}
