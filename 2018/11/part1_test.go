package main

import (
	"testing"
	"time"

	"github.com/kofalt/go-memoize"
	"github.com/stretchr/testify/assert"
)

func TestHundredsDigits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, get_hundreds_digit(1900))
	assert.Equal(6, get_hundreds_digit(632))
	assert.Equal(0, get_hundreds_digit(44))
}

func TestTotalPower(t *testing.T) {
	assert := assert.New(t)

	total_power_cache := memoize.NewMemoizer(time.Hour, time.Hour)

	// assert.Equal(4, calculate_total_power(33, 45, 18, 1))
	// assert.Equal(4, memoized_total_power(33, 45, 18, 1, total_power_cache))

	// assert.Equal(14, calculate_total_power(33, 45, 18, 2))
	// assert.Equal(14, memoized_total_power(33, 45, 18, 2, total_power_cache))

	assert.Equal(29, calculate_total_power(33, 45, 18, 3))
	assert.Equal(29, memoized_total_power(33, 45, 18, 3, total_power_cache))

	assert.Equal(1, calculate_total_power(33, 45, 18, 5))
	assert.Equal(1, memoized_total_power(33, 45, 18, 5, total_power_cache))

	x, y, power := compute_largest_total_power_square(18, 3, total_power_cache)
	assert.Equal(33, x)
	assert.Equal(45, y)
	assert.Equal(29, power)
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, fuel_cell_power_level(3, 5, 8))
	assert.Equal(-5, fuel_cell_power_level(122, 79, 57))
	assert.Equal(0, fuel_cell_power_level(217, 196, 39))
	assert.Equal(4, fuel_cell_power_level(101, 153, 71))
}
