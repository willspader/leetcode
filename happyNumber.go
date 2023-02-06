package main

import (
	"math"
)

func main() {
	isHappy(1)
}

func isHappy(n int) bool {
	var detectCycle map[int]bool = make(map[int]bool)

	var sum int = 0
	for n != 1 {
		sum = 0

		for n >= 1 {
			remainder := n % 10
			sum = sum + int(math.Pow(float64(remainder), float64(2)))
			n = n / 10
		}

		if detectCycle[sum] {
			return false
		}

		detectCycle[sum] = true
		n = sum
	}

	return true
}
