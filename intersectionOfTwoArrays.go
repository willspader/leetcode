package main

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int

	var i int = 0
	var j int = 0

	for i < len(nums1) && j < len(nums2) {
		for i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j] {
			i++
		}

		for j < len(nums2) && i < len(nums1) && nums2[j] < nums1[i] {
			j++
		}

		if i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
			if len(result) == 0 {
				result = append(result, nums1[i])
			} else if nums1[i] > nums1[i-1] {
				result = append(result, nums1[i])
			}
			i++
			j++
		}
	}

	return result
}
