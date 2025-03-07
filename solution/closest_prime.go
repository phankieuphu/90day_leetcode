package solution

import (
	"fmt"
	"math"
)

func closestPrimes(left int, right int) []int {
	result := []int{-1, -1}
	first, last := 0, 0
  min := right - left;
	for i := right; i >  left; i-- {
		if isPrime(i) == true {
      if last != 0 && last - first < min {
        min = last - first
        first = i
      }
      last = i
      fmt.Println("Last", last)
			// if first != 0 && i-first == 2 {
			// 	last = i
			// 	return []int{first, last}
			// }
			 
		}
	}
  if first != 0  && last  != 0{
    return []int{first,last}
  }
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func ClosestPrimes() {
  fmt.Println(  isPrime(25))  
  // listTestCase := [][]int{{4,6},{10,19},{19,31}}
  // for _,v := range listTestCase {
  //   result := closestPrimes(v[0],v[1])
  //   fmt.Println("Result", result)

  // } 
  return
}
