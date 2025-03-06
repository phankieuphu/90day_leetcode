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
  // if sliceMerged[0] != 1 {
  //   repeated = 1;
  // }
	for i := 0; i < len(sliceMerged); i++ {
    if i+1 != sliceMerged[i] {
      repeated = i
    }
    if i + 1 > sliceMerged[i] {
      duplicate = i
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
	input := [][]int{{2,2},{3,4}}
	result := findMissingAndRepeatedValues(input)
	fmt.Println("sliceMerged", result)
}
