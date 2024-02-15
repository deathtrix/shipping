package packing

import (
	"sort"
)

// Calculates packs to be sent for items
func Calculate(packSizes []int, items int) []int {
	// sort packSizes ascending
	sort.Ints(packSizes)

	// starting in reverse packSizes order, it fills packages
	numPacks := make([]int, len(packSizes))
	for i := len(packSizes) - 1; i >= 0; i-- {
		numPacks[i] = items / packSizes[i]
		items = items - numPacks[i]*packSizes[i]
	}

	// if items remain, go through packSizes ascending and check
	// if it can compact lower pack sizes into a bigger one
	if items > 0 {
		newPack := 1
		for i := 0; i < len(packSizes)-1; i++ {
			if (numPacks[i]+newPack)*packSizes[i] >= packSizes[i+1] {
				numPacks[i+1]++
				numPacks[i] = 0
			} else {
				numPacks[i] += newPack
			}
			newPack = 0
		}
	}

	return numPacks
}
