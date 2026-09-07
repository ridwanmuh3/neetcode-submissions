func missingNumber(nums []int) int {
	n := len(nums)
	xorr := n
	for i := range n {
		xorr ^= i ^ nums[i]
	}
	return xorr
}
