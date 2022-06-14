/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	map1 := map[int]int{}
	for i, x := range nums {
		if p, ok := map1[target-x]; ok {
			return []int{p, i}
		}
		map1[x] = i
	}
	return nil
}

// @lc code=end

