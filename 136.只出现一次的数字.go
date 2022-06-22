/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	res := 0
	// 一个数与0异或等于这个数，一个数与自身异或等于0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// @lc code=end

