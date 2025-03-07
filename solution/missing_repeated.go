package solution

import (
	"fmt"
	"sort"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	sliceMerged := []int{}
	for _, v := range grid {
		sliceMerged = append(sliceMerged, v...)
	}
	sort.Ints(sliceMerged)
	duplicate, repeated := 0, -1
	fmt.Println(sliceMerged)
	if sliceMerged[0] != 1 {
		repeated = 1
	}
	for i := 1; i < len(sliceMerged); i++ {
		if sliceMerged[i] == sliceMerged[i-1] {
			duplicate = sliceMerged[i]
		}
		if sliceMerged[i] > sliceMerged[i-1]+1 {
			repeated = sliceMerged[i-1] + 1
		}
	}
	if repeated == -1 {
		repeated = sliceMerged[len(sliceMerged)-1] + 1
	}
	return []int{duplicate, repeated}
}

// Input: grid = [[1,3],[2,2]]
// Output: [2,4]
// Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4]
func FindMissing() {
	input := [][]int{{2, 2}, {3, 4}}
	result := findMissingAndRepeatedValues(input)
	fmt.Println("sliceMerged", result)
}
