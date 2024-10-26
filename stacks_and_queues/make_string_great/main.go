package main

import (
	"fmt"
	"strings"
)

/*
Given a string s of lower and upper case English letters.

A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:

0 <= i <= s.length - 2
s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.

Return the string after making it good. The answer is guaranteed to be unique under the given constraints.

Notice that an empty string is also good.



Example 1:

Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
*/
func main() {
	s := "abBAcC"
	fmt.Println(makeGood(s))
}

func makeGood(s string) string {
	if s == "" {
		return s
	}

	var stack []string

	for _, symbol := range strings.Split(s, "") {
		if adjacent(peek(stack), symbol) {
			stack = pop(stack)
		} else {
			stack = push(stack, symbol)
		}
	}

	return strings.Join(stack, "")
}

func adjacent(s1 string, s2 string) bool {
	return (strings.ToUpper(s1) == strings.ToUpper(s2)) && (s1 != s2)
}

func peek(stack []string) string {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return ""
}

func pop(stack []string) []string {
	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}
	return stack
}

func push(stack []string, dir string) []string {
	return append(stack, dir)
}
