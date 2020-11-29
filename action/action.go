package action

import (
	"fmt"
)

type ActionType int

const (
	ActionTypeUnknown ActionType = iota
	ActionTypePickup
	ActionTypeAdded
	ActionTypeSelectShip
	ActionTypeHyper
	ActionTypeLevelUp
	ActionTypeScrap
	ActionTypeAttack
	ActionTypeRepair
	ActionTypeSelfRep
)

type Action interface {
	fmt.Stringer
	ActionType() ActionType
	Tick() int
}
