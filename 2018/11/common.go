package main

import (
	"fmt"
	"math"
	"time"

	"github.com/kofalt/go-memoize"
)

func get_hundreds_digit(num int) int {
	sub_thousand := num % 1000
	hundreds_digit := sub_thousand / 100

	return hundreds_digit
}

func fuel_cell_power_level(x, y int, serial int) int {
	rack_id := x + 10
	power_level := rack_id * y
	power_level += serial
	power_level *= rack_id

	power_level = get_hundreds_digit(power_level)
	power_level -= 5

	return power_level
}

func calculate_total_power(top_left_x, top_left_y int, serial int, square_size int) int {
	total_power := 0

	bottom_right_x := top_left_x + square_size
	bottom_right_y := top_left_y + square_size

	for x := top_left_x; x < bottom_right_x; x++ {
		for y := top_left_y; y < bottom_right_y; y++ {
			delta := fuel_cell_power_level(x, y, serial)

			total_power += delta
			// fmt.Println("addzing", x, y, delta, total_power)
		}
	}
	return total_power
}

func memoized_total_power(top_left_x, top_left_y int, serial int, square_size int, cache *memoize.Memoizer) int {
	// fmt.Println("memoized_total_power", top_left_x, top_left_y, serial, square_size)

	switch square_size {
	case 1:
		return fuel_cell_power_level(top_left_x, top_left_y, serial)
	default:
		key := fmt.Sprintf("%d-%dx%d", top_left_x, top_left_y, square_size-1)

		var total_power int
		total_power_obj, _, _ := cache.Memoize(key, func() (interface{}, error) {
			// fmt.Println("running from memoize")
			return memoized_total_power(top_left_x, top_left_y, serial, square_size-1, cache), nil
		})
		total_power = total_power_obj.(int)

		// fmt.Println("total_power", key, square_size, total_power)

		bottom_right_x := top_left_x + square_size
		bottom_right_y := top_left_y + square_size

		right_hand_x := bottom_right_x - 1
		bottom_y := bottom_right_y - 1

		for x := top_left_x; x < bottom_right_x; x++ {
			delta := fuel_cell_power_level(x, bottom_y, serial)
			// fmt.Println("adding", x, bottom_y, delta)
			total_power += delta
		}

		for y := top_left_y; y < bottom_right_y-1; y++ {
			delta := fuel_cell_power_level(right_hand_x, y, serial)
			// fmt.Println("adding", right_hand_x, y, delta)
			total_power += delta
		}

		return total_power
	}
}

func compute_largest_total_power_square(serial int, square_size int, cache *memoize.Memoizer) (x int, y int, power int) {
	highestPower := math.MinInt32
	highestPowerX := -1
	highestPowerY := -1

	highest_xy := 300 - square_size + 1

	for x := 0; x <= highest_xy; x++ {
		for y := 0; y <= highest_xy; y++ {
			// power := calculate_total_power(x, y, serial, square_size)

			power := memoized_total_power(x, y, serial, square_size, cache)

			if power > highestPower {
				highestPower = power
				highestPowerX = x
				highestPowerY = y
			}
		}
	}

	return highestPowerX, highestPowerY, highestPower
}

func computePart2(serial int) (x, y int, power int, size int) {
	highestPower := math.MinInt32
	highestPowerSize := 0
	highestPowerX := -1
	highestPowerY := -1

	total_power_cache := memoize.NewMemoizer(time.Hour, time.Hour)

	for size := 1; size <= 300; size++ {
		x, y, power := compute_largest_total_power_square(serial, size, total_power_cache)
		if power > highestPower {
			highestPower = power
			highestPowerSize = size
			highestPowerX = x
			highestPowerY = y
		}
	}

	return highestPowerX, highestPowerY, highestPower, highestPowerSize
}
