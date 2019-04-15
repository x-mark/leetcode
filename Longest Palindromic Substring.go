/*

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

*/

package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(longestPalindrome("aacca"))
// }

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	subBegin, subEnd := 0, 0
	for pos := 0; pos < 2*(len(s)-1); pos++ {
		begin, end := 0, 0
		if pos%2 == 0 {
			if pos == 0 {
				begin = 0
				end = 0
			} else {
				begin = pos/2 - 1
				end = pos/2 + 1
			}
		} else {
			begin = (pos - 1) / 2
			end = (pos + 1) / 2
		}

		for s[end] == s[begin] {
			if subEnd-subBegin < end-begin {
				subBegin = begin
				subEnd = end
			}
			if begin > 0 && end < len(s)-1 {
				begin--
				end++
			} else {
				break
			}
		}
	}
	return s[subBegin : subEnd+1]
}
