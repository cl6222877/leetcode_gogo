/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	length := len(nums)
	var maxIndex int
	// 不用遍历到最后一个
	for i := 0; i < length-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		// 如果此时最远小于i 或 遇到0
		if maxIndex <= i {
			return false
		}

	}
	return maxIndex >= length-1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

