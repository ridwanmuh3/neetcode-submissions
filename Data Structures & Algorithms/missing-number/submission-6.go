func missingNumber(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	n := len(nums)
	for i := range n + 1 {
		if _, exists := set[i]; !exists {
			return i
		}
	}
	return -1
}
