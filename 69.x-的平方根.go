/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	ans := -1
	left, right := 0, x
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle <= x {
			ans = middle
			left = middle + 1

		} else {
			right = middle - 1
		}
	}

	return ans

}

// @lc code=end

