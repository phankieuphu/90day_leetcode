package solution

import "fmt"

func twoSum(numbers []int, target int) []int {
	result := []int{}

//	mid := len(numbers) / 2
	start := 0
	end := len(numbers) -1
	// if target <= mid {
	// 	end = mid
	// }

	for start <= end {
		count := numbers[start] + numbers[end]
    fmt.Println(count)
		if count == target {
			return []int{start + 1, end + 1}
		}
		if count < target {
			start++
		}
		if count > target {
			count--
		}
	}
  return result
}


func TwoSum(){
  input := []int{-1000,-1,0,1}
  target := 1;
  twoSum(input,target)
}