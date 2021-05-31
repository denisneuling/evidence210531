package main

import (
	"fmt"
	"sort"
)

// Merge, sort and eliminate overlapping intervals
func Merge(inputIntervals [][]int) ([][]int, error) {
	intervals := make(map[int]int)

	for _, inputElement := range inputIntervals {
		if len(inputElement) != 2 {
			return nil, fmt.Errorf("interval must contain exactly 2 elements, got %v", inputElement)
		}

		min := inputElement[0]
		max := inputElement[1]

		if len(intervals) == 0 {
			intervals[min] = max
			continue
		}

		intervalToBeAdded := true
		for currentMin, currentMax := range intervals {

			// already covered, skip
			if inBetween(min, currentMin, currentMax) && inBetween(max, currentMin, currentMax) {
				intervalToBeAdded = false
				break
			}

			// min in boundaries but max needs adjustment
			if inBetween(min, currentMin, currentMax) && max > currentMax {
				currentMax = max
				intervals[currentMin] = currentMax
				intervalToBeAdded = false
			}

			// max in boundaries but min needs adjustment
			if inBetween(max, currentMin, currentMax) && min < currentMin {
				intervals[min] = currentMax
				// delete the obsolete/old interval
				delete(intervals, currentMin)
				currentMin = min
				intervalToBeAdded = false
			}
		}

		if intervalToBeAdded {
			intervals[min] = max
		}
	}

	// convert map to sorted array of intervals
	result := mapToArrayOfInts(intervals)

	return result, nil
}

// check whether given value is within a specific boundary
func inBetween(i int, min int, max int) bool {
	return (i >= min) && (i <= max)
}

// convert a map of ints to slice of int slices
func mapToArrayOfInts(m map[int]int) [][]int {
	result := make([][]int, 0)

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, key := range keys {
		result = append(result, []int{key, m[key]})
	}

	return result
}
