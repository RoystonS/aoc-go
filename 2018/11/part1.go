package main

import (
	"aoccommon"
	"fmt"
	"strconv"
	"time"

	"github.com/kofalt/go-memoize"
)

func part1() {
	lines, err := aoccommon.ReadLines("input")
	aoccommon.CheckError(err)

	serial, _ := strconv.Atoi(lines[0])
	total_power_cache := memoize.NewMemoizer(time.Hour, time.Hour)

	x, y, _ := compute_largest_total_power_square(serial, 3, total_power_cache)

	fmt.Printf("%d,%d\n", x, y)
}
