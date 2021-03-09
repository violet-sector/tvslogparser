package action_test

import (
	"testing"

	"github.com/AlexCrane/tvslogparser/action"
	"github.com/AlexCrane/tvslogparser/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Round42_Attack(t *testing.T) {
	me := common.NewPlayer("Cranar272")
	me.SelectShip("Fanged Fighter")

	attackOut := []string{"2", "Cranar272", "Rico", "Cranar272 attacked Rico", "-41", "1809"}
	att, err := action.AttackFromCSVRecord(attackOut, me)
	require.NoError(t, err)
	t.Log(att)
	assert.True(t, att.Outgoing)
	assert.Equal(t, "Rico", att.EnemyPilot)
	assert.Equal(t, 41, att.TargetDamage)
	assert.Equal(t, 1809, att.TargetHitpointsRemaining)
}

func Test_Round43_Attack(t *testing.T) {
	me := common.NewPlayer("Red October")
	me.SelectShip("Cloud of Death")

	attackOutRoids := []string{"2", "Red October", "Guderian", "Red October attacked Guderian", "-16", "1010", "-12", "1628"}
	att, err := action.AttackFromCSVRecord(attackOutRoids, me)
	require.NoError(t, err)
	t.Log(att)
	assert.True(t, att.Outgoing)
	assert.Equal(t, "Guderian", att.EnemyPilot)
	assert.Equal(t, 16, att.TargetDamage)
	assert.Equal(t, 1010, att.TargetHitpointsRemaining)
	assert.Equal(t, 12, att.PlayerDamageTaken)
	assert.Equal(t, 1628, att.PlayerHitpointsRemaining)

	attackIn := []string{"168", "not the LC", "Red October", "not the LC attacked Red October", "-321", "11953", "N/A", "N/A"}
	att, err = action.AttackFromCSVRecord(attackIn, me)
	require.NoError(t, err)
	t.Log(att)
	assert.False(t, att.Outgoing)
	assert.Equal(t, "not the LC", att.EnemyPilot)
	assert.Equal(t, 321, att.TargetDamage)
	assert.Equal(t, 11953, att.TargetHitpointsRemaining)
	assert.Equal(t, 0, att.PlayerDamageTaken)
	assert.Equal(t, 0, att.PlayerHitpointsRemaining)
}

func Test_Kill(t *testing.T) {
	// TODO
}

func Test_Death(t *testing.T) {
	// TODO
}

func Test_SuicidalAttack(t *testing.T) {
	// TODO
}
