package packing

import (
	"math"
	"sort"
	"sync"
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

func CalculateConc(packSizes []int, items int) []int {
	minItems := math.MaxInt
	minPacks := math.MaxInt
	numPacks := []int{}
	var wg sync.WaitGroup
	var mut sync.Mutex

	for i := len(packSizes) - 1; i > 0; i-- {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup, mut *sync.Mutex) {
			defer wg.Done()
			cNumPacks, cItems, cPacks := posCalculate(packSizes, items, i)
			mut.Lock()
			if cPacks < minPacks || (cItems >= items && cItems < minItems) {
				minPacks = cPacks
				minItems = cItems
				numPacks = cNumPacks
			}
			mut.Unlock()
		}(i, &wg, &mut)
	}
	wg.Wait()

	return numPacks
}

func posCalculate(packSizes []int, items, pos int) ([]int, int, int) {
	// sort packSizes ascending
	sort.Ints(packSizes)

	cntItems := 0
	cntPacks := 0
	// starting in reverse packSizes order, it fills packages
	numPacks := make([]int, len(packSizes))
	for i := pos; i >= 0; i-- {
		numPacks[i] = items / packSizes[i]
		if numPacks[i] > 0 {
			items = items - numPacks[i]*packSizes[i]
			cntItems += numPacks[i] * packSizes[i]
			cntPacks += numPacks[i]
		}
	}

	// if items remain, go through packSizes ascending and check
	// if it can compact lower pack sizes into a bigger one
	if items > 0 {
		newPack := 1
		for i := 0; i < len(packSizes)-1; i++ {
			if (numPacks[i]+newPack)*packSizes[i] >= packSizes[i+1] {
				cntItems += packSizes[i]
				numPacks[i+1]++
				numPacks[i] = 0
			} else {
				if newPack > 0 {
					cntItems += packSizes[i]
					numPacks[i] += newPack
					cntPacks++
				}
			}
			newPack = 0
		}
	}

	return numPacks, cntItems, cntPacks
}
