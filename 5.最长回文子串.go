/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

@lc code=start
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	start, end := 0, 0
	for i, _ := range s {
		// for i := 0; i < len(s); i++ {
		left1, right1 := extendPa(s, i, i)
		left2, right2 := extendPa(s, i, i+1)

		length1 := right1 - left1
		length2 := right2 - left2
		lenLong := end - start
		if length1 > lenLong {
			start = left1
			end = right1
		}
		if length2 > lenLong {
			start = left2
			end = right2
		}

	}
	return s[start : end+1]
}

func extendPa(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}


// @lc code=end


