func isAnagram(s string, t string) bool {
	occuredCharS := map[rune]int{}
	occuredCharT := map[rune]int{}
	for _, charS := range s {
		occuredCharS[charS]++
	}
	for _, charT := range t {
		occuredCharT[charT]++
	}

	isValidAnagram := true
	for _, charT := range t {
		countCharT, _ := occuredCharT[charT]
		countCharS, _ := occuredCharS[charT]
		if countCharT != countCharS {
			isValidAnagram = false
		}
	}

	return isValidAnagram && len(s) == len(t) 
}
