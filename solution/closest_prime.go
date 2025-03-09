package solution

import (
	"fmt"
	"math"
)

func closestPrimes(left int, right int) []int {
  result := []int{-1,-1}
  first, last := 0,0

  holdResult := []int{}
  for i := left; i <= right;i++ {
    if isPrime(i) == true {
      holdResult = append(holdResult, i)
    }
  }
  min := right - left;
  for i := len(holdResult) - 1;i > 0; i-- {
    if min >= holdResult[i] - holdResult[i-1] {
      last = holdResult[i]
      first = holdResult[i-1]
      min = holdResult[i] -  holdResult[i-1]
    }
  } 
  if first != 0 && last != 0 {
    return []int{first,last}
  }
  return result
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func ClosestPrimes() {

   listTestCase := [][]int{{4,6},{10,19},{19,31}}

  for _,v := range listTestCase {
    result := closestPrimes(v[0],v[1])
    fmt.Println("Result", result)
  } 
  return
}
