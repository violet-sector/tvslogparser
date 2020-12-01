package action_test

import (
	"encoding/csv"
	"strings"
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

func Test_Pickup(t *testing.T) {
	var pickupTests = []struct {
		csvLine            string
		expectedPickupType action.PickupType
	}{
		{`18,Cranar272,,"Pickup (Gate intel: T18 The Uncharted Sector->The Asteroid Field T18)",0,0`, action.PTGateIntel},
		{`1,Cranar272,,"Pickup (None)",0,0`, action.PTNone},
		{`20,Cranar272,,"Pickup (150 points)",0,0`, action.PTPoints150},
		{`23,Cranar272,,"Pickup (100 points)",0,0`, action.PTPoints100},
		{`24,Cranar272,,"Pickup (Free attack)",0,0`, action.PTFreeAttack},
		{`26,Cranar272,,"Pickup (300 points)",0,0`, action.PTPoints300},
		{`28,Cranar272,,"Pickup (Triple attack)",0,0`, action.PTTripleAttack},
		{`31,Cranar272,,"Pickup (250 points)",0,0`, action.PTPoints250},
		{`38,Cranar272,,"Pickup (Half max repair)",0,0`, action.PTHalfMaxRepair},
		{`42,Cranar272,,"Pickup (-250 hitpoints)",0,0`, action.PTExplosive250},
		{`46,Cranar272,,"Pickup (Invulnerability)",0,0`, action.PTInvulnerability},
		{`65,Cranar272,,"Pickup (Random hyper gate)",0,0`, action.PTRandomHyperGate},
		{`69,Cranar272,,"Pickup (-450 hitpoints)",0,0`, action.PTExplosive450},
		{`104,Cranar272,,"Pickup (Decoy)",0,0`, action.PTDecoy},
		{`187,Cranar272,,"Pickup (-50 hitpoints)",0,0`, action.PTExplosive50},
		{`147,Cranar272,,"Pickup (200 hitpoints)",0,0`, action.PTRepair200},
		{`194,Cranar272,,"Pickup (150 hitpoints)",0,0`, action.PTRepair150},
		{`249,Cranar272,,"Pickup (300 hitpoints)",0,0`, action.PTRepair300},
	}
	for _, tt := range pickupTests {
		logReader := csv.NewReader(strings.NewReader(tt.csvLine))
		csvSlice, err := logReader.Read()
		require.NoError(t, err)
		pickup, err := action.PickupFromCSVRecord(csvSlice)
		require.NoError(t, err)
		assert.Equal(t, tt.expectedPickupType, pickup.Type)
	}
}
