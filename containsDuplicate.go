package main

func containsDuplicate(nums []int) bool {
	var contains map[int]bool = make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if contains[nums[i]] {
			return true
		} else {
			contains[nums[i]] = true
		}
	}
	return false
}
