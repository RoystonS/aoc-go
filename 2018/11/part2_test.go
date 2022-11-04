package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	x, y, power, square_size := computePart2(18)
	assert.Equal(90, x)
	assert.Equal(269, y)
	assert.Equal(113, power)
	assert.Equal(16, square_size)

	x, y, power, square_size = computePart2(42)
	assert.Equal(232, x)
	assert.Equal(251, y)
	assert.Equal(119, power)
	assert.Equal(12, square_size)
}
