package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("findMedianSortedArrays(nums1, nums2): %v\n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	// f := int(math.Round(float64((n + m) / 2)))
	f := n + m
	fmt.Printf("f: %v\n", f)

	key1 := 0
	key2 := 0

	c := make([]int, f)
	switch {
	case n == 0:
		c = nums2
	case m == 0:
		c = nums1
	default:

		for count := 0; count < f/2+1; count++ {
			switch {
			case key1 == n:
				c[count] = nums2[key2]
				key2++
			case key2 == m:
				c[count] = nums1[key1]
				key1++
			default:
				if nums1[key1] <= nums2[key2] {
					c[count] = nums1[key1]
					// fmt.Printf("key1: %v\n", key1)
					key1++
				} else {
					c[count] = nums2[key2]
					// fmt.Printf("key2: %v\n", key2)
					key2++

				}
			}
			// fmt.Printf("c: %v\n", c)
		}
	}
	length := len(c)
	fmt.Printf("length: %v\n", length)
	if len(c) == 1 {
		return float64(c[0])
	}
	// fmt.Printf("c[2]: %v\n", c[2])
	a := float64(c[(n+m)/2-1])
	b := float64(c[(n+m)/2])
	d := (a + b) / 2
	// fmt.Printf("d: %v\n", d)
	if (n+m)%2 == 1 {
		return b

	} else {
		return d
	}
}
