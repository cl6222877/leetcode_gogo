/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		digits[i]++
		// 10进制
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	// 如果加起来全部为0
	digits = make([]int, length+1)
	digits[0] = 1
	return digits
}

// @lc code=end
