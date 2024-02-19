package packing

import (
	"math"
	"sort"
)

// Calculates packs to be sent for items
func Calculate(packSizes []int, items int) []int {
	// sort packSizes ascending
	sort.Ints(packSizes)

	minItems := math.MaxInt
	minPacks := math.MaxInt
	numPacks := []int{}

	// calculate for each position in pack size
	for j := len(packSizes) - 1; j > 0; j-- {
		cntItems := 0
		cntPacks := 0
		cItems := items
		cntNumPacks := make([]int, len(packSizes))
		// starting in reverse packSizes order, it fills packages
		for i := j; i >= 0; i-- {
			cntNumPacks[i] = cItems / packSizes[i]
			if cntNumPacks[i] > 0 {
				cItems = cItems - cntNumPacks[i]*packSizes[i]
				cntItems += cntNumPacks[i] * packSizes[i]
				cntPacks += cntNumPacks[i]
			}
		}

		// if items remain, go through packSizes ascending and check
		// if it can compact lower pack sizes into a bigger one
		if cItems > 0 {
			newPack := 1
			for i := 0; i < len(packSizes)-1; i++ {
				if (cntNumPacks[i]+newPack)*packSizes[i] >= packSizes[i+1] {
					cntItems += packSizes[i]
					cntNumPacks[i+1]++
					cntNumPacks[i] = 0
				} else {
					if newPack > 0 {
						cntItems += packSizes[i]
						cntNumPacks[i] += newPack
						cntPacks++
					}
				}
				newPack = 0
			}
		}

		// find solution with minimum items and minimum packs from solutions
		// found for each position
		if cntPacks < minPacks || (cntItems >= items && cntItems < minItems) {
			minPacks = cntPacks
			minItems = cntItems
			numPacks = cntNumPacks
		}
	}

	return numPacks
}
