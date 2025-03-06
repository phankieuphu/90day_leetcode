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
  duplicate,repeated := 0,0
  for i:= 1 ; i < len(sliceMerged) -1;i++ {
    if sliceMerged[i] == sliceMerged[i-1] {
        duplicate = sliceMerged[i]
    }
    if sliceMerged[i] > sliceMerged[i-1] +1 {
      repeated = sliceMerged[i-1] + 1
    }
  }
	return []int{duplicate, repeated}
}

// Input: grid = [[1,3],[2,2]]
// Output: [2,4]
// Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4]
func FindMissing() {
	input := [][]int{{1, 3}, {2, 2}}
	result := findMissingAndRepeatedValues(input)
	fmt.Println("sliceMerged", result)
}
