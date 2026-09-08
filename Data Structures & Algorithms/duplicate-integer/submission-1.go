func hasDuplicate(nums []int) bool {
    duplicatedNums := map[int]bool{}
    for _, num := range nums {
        if _, ok := duplicatedNums[num]; ok {
            duplicatedNums[num] = true;
            return true
        }
        duplicatedNums[num] = false;
    }
    return false
}
