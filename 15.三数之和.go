/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
// 此方法使用leecode官方方法
func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for first := 0; first < length; first++ {
		// 过滤和上个数一样的值，避免重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 因为1 2都是从最小开始找，3就从最大开始找可以减小计算
		third := length - 1
		target := -1 * nums[first]
		for second := first + 1; second < length; second++ {
			// 过滤和上个数一样的值，避免重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 注意这里是while 因为third从大到校，基本上只会大于，小于等于说明没有直接跳过
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}

		}
	}
	return res
}

// @lc code=end

