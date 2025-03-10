package main

import (
	"fmt"
	"main/solution"
	"strings"
)

func main() {
	fmt.Println("Start programing")
	// solution.FindMissing()
	// solution.ClosestPrimes()
	//solution.ClosestPrimes()
	solution.MinimunSize()
}

func permute(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		return result
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	for i := 0; i < len(nums); i++ {
		holdSlice := []int{}
		for j := i; j < len(nums); j++ {
			holdSlice = append(holdSlice, nums[j])
		}
		result = append(result, holdSlice)
	}

	return result
}

func findSubstring(s string, words []string) []int {
	// sure word include in string
	validWords := []string{}
	for _, value := range words {
		if strings.Contains(s, value) == true {
			validWords = append(validWords, value)
		}
	}
	if len(validWords) == 0 {
		return []int{}
	}
	startWork := strings.Index(s, validWords[0])
	endWorkd := strings.Index(s, validWords[0])
	for index, _ := range validWords {
		whereSub := findAllIndex(s, validWords[index])
		if len(whereSub) != 0 {
			if whereSub[0] < startWork {
				startWork = whereSub[0]
			} else {
				endWorkd = max(endWorkd, whereSub[len(whereSub)-1])
			}
			//	startWork = min(startWork, whereSub[0])

		}
		// startWork = min(startWork, strings.Index(s, validWords[index]))
		// endWorkd = max(endWorkd, strings.Index(s, validWords[index]))

	}
	return []int{startWork, endWorkd}
}
func findAllIndex(str, subString string) []int {
	var indexs []int
	start := 0
	for {
		index := strings.Index(str[start:], subString)
		if index == -1 {
			break
		}
		indexs = append(indexs, start+index)
		start += index + len(subString)
	}
	//	fmt.Println(indexs)
	return indexs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
