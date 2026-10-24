func lengthOfLongestSubstring(s string) int {
    left := 0
    right := 0
    dic := make(map[byte]int)
    ret := 0
    for left <= right && right < len(s) {
        for { 
            if _, ok := dic[s[right]];ok {
                dic[s[left]] = dic[s[left]] - 1
                if dic[s[left]] == 0 {
                    delete(dic, s[left])
                }
                left += 1
            } else {
                break
            }
        }
        dic[s[right]] = 1
        ret = max(ret, right - left + 1)
        right += 1
    }
    return ret
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}