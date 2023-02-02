package main

import (
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)

	var size int = len(nums)

	if size < 3 {
		return nums[size-1]
	}

	var bigger int = nums[size-1]
	var idx int = size - 2
	var changes int = 0

	for changes < 2 && idx >= 0 {
		if nums[idx] != bigger {
			bigger = nums[idx]
			changes++
		}
		idx--
	}

	if changes < 2 {
		return nums[size-1]
	}

	return nums[idx+1]
}
