package main

import (
	"fmt"
	"os"
	"sync"

	"aoccommon"
)

func all_tiles_count_one(tiles [1000][1000]int, claim Claim) bool {
	for y := claim.top; y < claim.top+claim.height; y++ {
		for x := claim.left; x < claim.left+claim.width; x++ {
			if tiles[x][y] != 1 {
				return false
			}
		}
	}
	return true
}

func part2() {
	claims, err := read_claims("input")
	aoccommon.CheckError(err)

	tiles := apply_claims(claims)

	var wg sync.WaitGroup

	// Find the claim where all of its tiles are count 1
	for _, claim := range claims {
		claim := claim

		wg.Add(1)

		go func() {
			defer wg.Done()
			if all_tiles_count_one(tiles, claim) {
				fmt.Println(claim.id)
				os.Exit(0)
			}
		}()
	}

	wg.Wait()
}
