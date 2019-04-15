/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
package main

// import (
// 	"fmt"
// )

// func main() {
// 	nums1 := []int{1, 4}
// 	nums2 := []int{2}
// 	fmt.Print(findMedianSortedArrays(nums1, nums2))
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	iMax := len(nums1)
	jMax := len(nums2)
	p, q := 0, 0
	if (iMax+jMax)%2 == 0 {
		q = (iMax + jMax) / 2
		p = q - 1
	} else {
		q = (iMax + jMax) / 2
		p = q
	}

	index := 0
	sum := 0
	value := 0
	for {
		switch {
		case i < iMax && j < jMax:
			if nums1[i] < nums2[j] {
				value = nums1[i]
				i++
			} else {
				value = nums2[j]
				j++
			}
		case i < iMax && j >= jMax:
			value = nums1[i]
			i++
		case i >= iMax && j < jMax:
			value = nums2[j]
			j++
		}

		if index == p {
			sum += value
		}

		if index == q {
			sum += value
			break
		}

		index++
	}

	return float64(sum) / 2.0

}
