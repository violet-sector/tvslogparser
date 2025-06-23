package common_test

import (
	"testing"

	"github.com/AlexCrane/tvslogparser/common"
	"github.com/stretchr/testify/assert"
)

func Test_Ship(t *testing.T) {
	testCruiser := &common.ShipLevel{
		Level: 6,
	}

	assert.True(t, testCruiser.IsCruiser())

	testLevelFive := &common.ShipLevel{
		Level: 5,
	}

	assert.False(t, testLevelFive.IsCruiser())
}
