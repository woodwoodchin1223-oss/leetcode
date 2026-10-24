func twoSum(nums []int, target int) []int {
    ret := make([]int, 2)
    for i, n := range nums {
        for j, m := range nums[i + 1:] {
            if n + m == target {
                ret[0] = i
                ret[1] = i + 1 + j
                return ret
            }
        }
    }
    return ret
}