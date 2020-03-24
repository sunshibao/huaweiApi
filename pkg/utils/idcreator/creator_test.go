package idcreator

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCreator(t *testing.T) {

	currentLogIdCreator := idCreator

	InitCreator(uint16(rand.Int31()))

	assert.NotEqual(t, currentLogIdCreator, idCreator)
}

func TestNextID(t *testing.T) {

	id1 := NextID()
	assert.Greater(t, id1, uint64(0))

	id2 := NextID()
	assert.Greater(t, id2, uint64(0))

	assert.NotEqual(t, id1, id2)
}

func TestNextString(t *testing.T) {
	str1 := NextString()
	assert.NotEqual(t, str1, "0")

	str2 := NextString()
	assert.NotEqual(t, str2, "0")

	assert.NotEqual(t, str1, str2)
}
