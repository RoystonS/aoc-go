package main

import (
	"regexp"
	"strconv"

	"aoccommon"
)

type Claim struct {
	id int

	top    int
	left   int
	width  int
	height int
}

func parse_line(line string) Claim {
	// line is of the form #1226 @ 679,914: 10x21
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	res := re.FindAllStringSubmatch(line, -1)

	id, err := strconv.Atoi(res[0][1])
	aoccommon.CheckError(err)
	left, err := strconv.Atoi(res[0][2])
	aoccommon.CheckError(err)
	top, err := strconv.Atoi(res[0][3])
	aoccommon.CheckError(err)
	w, err := strconv.Atoi(res[0][4])
	aoccommon.CheckError(err)
	h, err := strconv.Atoi(res[0][5])
	aoccommon.CheckError(err)

	return Claim{
		id:  id,
		top: top, left: left,
		width: w, height: h}
}

func read_claims(filename string) (claims []Claim, err error) {
	lines, err := aoccommon.ReadLines(filename)
	if err == nil {

		for _, line := range lines {
			claims = append(claims, parse_line(line))
		}
	}

	return claims, err
}

// Applies a series of claims to fabric, returning
// the number of times each square was claimed
func apply_claims(claims []Claim) (tiles [1000][1000]int) {
	// Apply all the claims to the fabric
	for _, claim := range claims {
		for y := claim.top; y < claim.top+claim.height; y++ {
			for x := claim.left; x < claim.left+claim.width; x++ {
				tiles[x][y]++
			}
		}
	}

	return tiles
}
