// this is hard problem
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".

// (()())()()())) -> 12

package solution

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	result := 0
	testItem := []int{}
	for i := 0; i < len(s); i++ {
		testItem = append(testItem, int(s[i]))
	}
	fmt.Println(testItem)

	return result
}

func matchString(a, b string) bool {
	if a == "(" && b == ")" {
		return true
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func LongestValid() {
	input := "(()())()()()))"
	longestValidParentheses(input)
}
