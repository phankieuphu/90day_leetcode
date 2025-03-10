package solution

import (
	"fmt"
	"sort"
)

// 209. Minimum Size Subarray Sum
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2

func minSubArrayLen(target int, nums []int) int {
	// result := []int{}
	sort.Ints(nums)
  fmt.Println(nums)
	if target < nums[0] {

		return 0
	}
	sumV := 0
	for _, v := range nums {
		if target >= v {
			return 1
		}
		sumV += v
	}
	// out range
	if sumV < target {
		return 0
	}
	if sumV == target {
		return len(nums)
	}

	start := 0
	end := len(nums) - 1
	// mid := len(nums) /2

	// if target > nums[len(nums)-1] {
	//   start = len(nums) -1;
	//   end = 0
	// }
	// if target < nums[mid]  {
	//   start = 0
	//   end = mid
	//
	minV := 0
  holdValue := 0

	for start < end {
		if nums[end]+nums[end-1]+holdValue >= target {
			fmt.Println("holdValue", nums[end-1 :])
			return len(nums[end-1 :])
		}
    holdValue += nums[end]
		end--
	}

	return minV
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func MinimunSize() {
	target := 213
	nums := []int{12,28,83,4,25,26,25,2,25,25,25,12}
	minSubArrayLen(target, nums)
}
