/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Print(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	var window []rune
	max := 0
	for _, r := range s {
		for i, vr := range window {
			if vr == r {
				window = window[i+1:]
				break
			}
		}

		window = append(window, r)
		if max < len(window) {
			max = len(window)
		}
	}

	return max
}

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) < 2 {
// 		return len(s)
// 	}
// 	sMap := make(map[rune][]int)
// 	start, max := 0, 1
// 	for i, r := range s {
// 		list, ok := sMap[r]
// 		if ok && start < list[len(list)-1]+1 {
// 			start = list[len(list)-1] + 1
// 		}

// 		sMap[r] = append(list, i)

// 		if max < i-start+1 {
// 			max = i - start + 1
// 		}
// 	}

// 	return max
// }
