package main

func findTheDifference(s string, t string) byte {
	var findCharAdded int = 0

	var i int = 0
	for i < len(s) {
		findCharAdded += int(s[i])
		findCharAdded -= int(t[i])
		i++
	}

	findCharAdded -= int(t[i])

	findCharAdded *= -1

	return byte(findCharAdded)
}
