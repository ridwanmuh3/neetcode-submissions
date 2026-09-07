func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}
	arr := make([][2]int, 0, len(freqs))
	for num, cnt := range freqs {
		arr = append(arr, [2]int{cnt, num})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	topK := make([]int, k)
	for i := range k {
		topK[i] = arr[i][1]
	}
	return topK
}
