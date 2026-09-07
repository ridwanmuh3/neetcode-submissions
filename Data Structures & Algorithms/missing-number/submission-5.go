func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := range n {
		if nums[i] != i {
			return i
		}
	}
	return n
}
