func singleNumber(nums []int) int {
	occurs := map[int]int{}
	for _, num := range nums {
		occurs[num]++
	}

	for num, count := range occurs {
		if count == 1 {
			return num
		}	
	}

	return 0
}
