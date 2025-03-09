package solution

import "fmt"

// Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// Output: 3
// constants
// n == gas.length == cost.length
// 1 <= n <= 105
// 0 <= gas[i], cost[i] <= 104

func canCompleteCircuit(gas []int, cost []int) int {

	if len(gas) != len(cost) || sum(gas) < sum(gas) {
    fmt.Println("sumCost")
		return -1
	}

  
	result := -1;
	addition := 0
	start, end := 0, len(gas)-1

	for start <= end {
		station := start
		if gas[station]+addition >= cost[station] {
			newGas := append(gas[station:], gas[:station]...)
			newCost := append(cost[station:], cost[:station]...)
      if isValid(newGas,newCost) == true {
        return start
      }
      start++
      // break

		} else {
			start++
			addition = 0
		}

	}
	return result
}

func isValid(gas,cost []int) bool {

  start,end := 0,len(gas)
  addition := 0;
  for start < end {
    if gas[start] + addition >= cost[start] {
      addition = gas[start] + addition - cost[start]
      start++
    }else {
      return false
    }
  }
  return true;
}
func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func CanCompleteCircuit() {
	gas, cost := []int{1,2,3,4,5,5,70}, []int{2,3,4,3,9,6,2}
	result := canCompleteCircuit(gas, cost)
	fmt.Println("Result", result)
}
