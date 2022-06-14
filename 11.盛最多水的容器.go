/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	len := len(height)
	ans := 0
	for left, right := 0, len-1; left < right; {
		leftNum := height[left]
		rightNum := height[right]
		length := right - left
		var minHeight int
		if leftNum > rightNum {
			minHeight = rightNum
			right--
		} else {
			minHeight = leftNum
			left++
		}
		area := length * minHeight
		if area > ans {
			ans = area
		}
	}
	return ans
}

// @lc code=end

