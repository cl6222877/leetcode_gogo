/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	// 暴力解法，KMP先不看了。
	length1 := len(haystack)
	length2 := len(needle)
	if length1 == 0 || length2 == 0 {
		return 0
	}
	if length2 > length1 {
		return -1
	}

	for i := 0; i <= length1-length2; i++ {
		// 首字母相等
		for j, k := i, 0; j < length1 && k < length2; {
			if haystack[j] != needle[k] {
				break
			}
			// 遍历到最后一位还相等
			if k == length2-1 {
				return i
			}
			j++
			k++
		}
	}

	return -1
}

// @lc code=end

