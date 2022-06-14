/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	// 第0个永远不会变，这里设置1
	slow := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow

}

// @lc code=end

