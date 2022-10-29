package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(32, runGame(9, 25))
	assert.Equal(8317, runGame(10, 1618))
	assert.Equal(146373, runGame(13, 7999))
	assert.Equal(54718, runGame(21, 6111))
	assert.Equal(37305, runGame(30, 5807))
}
