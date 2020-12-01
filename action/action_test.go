package action_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AlexCrane/tvslogparser/action"
)

func Test_LevelUp_V1(t *testing.T) {
	levelupV1 := []string{"7", "Cranar272", "", "Achieved level 2", "0", "0"}
	lvlUp, err := action.LevelUpFromCSVRecord(levelupV1)
	require.NoError(t, err)
	t.Log(lvlUp)
	assert.Equal(t, 2, lvlUp.Level.Level)
	assert.False(t, lvlUp.Level.IsCruiser())
	assert.Equal(t, false, lvlUp.Level.HasCombatPiloting)
}

func Test_LevelUp_V2(t *testing.T) {
	levelupV2 := []string{"7", "Cranar272", "", "Achieved level 2 (added 5 combat / 5 piloting)", "0", "0"}
	lvlUp, err := action.LevelUpFromCSVRecord(levelupV2)
	require.NoError(t, err)
	t.Log(lvlUp)
	assert.Equal(t, 2, lvlUp.Level.Level)
	assert.False(t, lvlUp.Level.IsCruiser())
	assert.Equal(t, true, lvlUp.Level.HasCombatPiloting)
	assert.Equal(t, 5, lvlUp.Level.CombatDelta)
	assert.Equal(t, 5, lvlUp.Level.PilotingDelta)
}

func Test_LevelUp_Cruiser(t *testing.T) {
	cruiser := []string{"205", "Cranar272", "", "Awarded cruiser carrier", "0", "0"}
	lvlUp, err := action.LevelUpFromCSVRecord(cruiser)
	require.NoError(t, err)
	t.Log(lvlUp)
	assert.Equal(t, 6, lvlUp.Level.Level)
	assert.True(t, lvlUp.Level.IsCruiser())
}
