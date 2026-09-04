func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, num := range nums {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
