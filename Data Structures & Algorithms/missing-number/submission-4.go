func missingNumber(nums []int) int {
	miss := make(map[int]bool)
	for _, num := range nums {
		miss[num] = true	
	}

	for i := range len(nums) + 1 {
		if _, exists := miss[i]; !exists {
			miss[i] = false
		}
	}

	for num, exists := range miss {
		if !exists {
			return num
		}
	}
	
	return -1
}
