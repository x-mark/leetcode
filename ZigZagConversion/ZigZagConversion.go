/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/

package leetcode

import (
	"bytes"
	// "fmt"
)

// func main() {
// 	fmt.Print(convert("AB", 1))
// }

func convert(s string, numRows int) string {
	arr := make(map[int][]byte)
	i, j := 0, 0
	positive := true
	for i = 0; i < len(s); i++ {
		arr[j] = append(arr[j], s[i])

		if positive {
			if j < numRows-1 {
				j++
			} else {
				positive = false
				if j > 0 {
					j--
				}
			}
		} else {
			if j > 0 {
				j--
			} else {
				positive = true
				if j < numRows-1 {
					j++
				}
			}
		}
	}

	var buffer bytes.Buffer
	for k := 0; k < numRows; k++ {
		buffer.Write(arr[k])
	}

	return buffer.String()
}
