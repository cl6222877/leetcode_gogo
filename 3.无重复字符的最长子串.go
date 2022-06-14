/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	uniqChar := map[byte]int{}
	rk, ans := -1, 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(uniqChar, s[i-1])
		}
		for rk+1 < length && uniqChar[s[rk+1]] == 0 {
			uniqChar[s[rk+1]]++
			rk++
		}
		index := rk - i + 1
		ans = max(&ans, &index)
	}
	return ans
}

func max(x, y *int) int {
	if *x > *y {
		return *x
	}
	return *y
}

// @lc code=end

