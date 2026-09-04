func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }

	countS, countT := map[rune]int{}, map[rune]int{}
	// alternate 
	for i, charS := range s {
		countS[charS]++
		countT[rune(t[i])]++
	}

	// original 
	// for _, charS := range s {
	// 	countS[charS]++
	// }
	// for _, charT := range t {
	// 	countT[charT]++
	// }


	for k, v := range countS {
		// original
		// countCharT, _ := countT[charT]
		// countCharS, _ := countS[charT]
		// if countCharT != countCharS {
		// 	return false
		// }

		// alternate
		if countT[k] != v {
			return false
		}
	}

	return true
}

// func sorted(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	sRunes, tRunes := []rune(s), []rune(t)
// 	sort.Slice(sRunes, func(i, j int) bool {
// 		return sRunes[i] < sRunes[j]
// 	})
// 	sort.Slice(sRunes, func(i, j int) bool {
// 		return tRunes[i] < tRunes[j]
//  	})

// 	for i := range sRunes {
// 		if sRunes[i] != tRunes[i] {
// 			return false
// 		}
// 	}
// 	return true
//  }
