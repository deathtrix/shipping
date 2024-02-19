package packing

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	nrPacks := []int{1, 250, 251, 501, 12001}
	nrPacksResults := [][]int{
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 1, 2},
	}

	for i, packs := range nrPacks {
		numPacks := Calculate(packSizes, packs)
		if !reflect.DeepEqual(numPacks, nrPacksResults[i]) {
			t.Fatalf("Wrong numpacks - %d", packs)
		}
	}
}

func TestCalculate2(t *testing.T) {
	packSizes := []int{3, 7, 13}
	nrPacks := []int{14}
	nrPacksResults := [][]int{
		{0, 2, 0},
	}

	for i, packs := range nrPacks {
		numPacks := Calculate(packSizes, packs)
		if !reflect.DeepEqual(numPacks, nrPacksResults[i]) {
			t.Fatalf("Wrong numpacks - %d", packs)
		}
	}
}

func BenchmarkCalculate(*testing.B) {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	packs := 251

	Calculate(packSizes, packs)
}
