func hasDuplicate(nums []int) bool {
    duplicatedNums := map[int]bool{}
    isDuplicate := false
    for _, num := range nums {
        if _, ok := duplicatedNums[num]; ok {
            duplicatedNums[num] = true;
            isDuplicate = true
        } else {
            duplicatedNums[num] = false;
        }
    }
    return isDuplicate
}
