package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var countMapS map[byte]int32 = make(map[byte]int32)
	var countMapT map[byte]int32 = make(map[byte]int32)

	for i := 0; i < len(s); i++ {
		if countMapS[s[i]] != 0 {
			countMapS[s[i]]++
		} else {
			countMapS[s[i]] = 1
		}

		if countMapT[t[i]] != 0 {
			countMapT[t[i]]++
		} else {
			countMapT[t[i]] = 1
		}
	}

	for key := range countMapS {
		if countMapS[key] != countMapT[key] {
			return false
		}
	}

	return true
}
