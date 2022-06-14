/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	best := math.MaxInt32

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		second, third := first+1, length-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			// 如果直接相等，返回之
			if sum == target {
				return sum
			}

			update(sum)

			// 更新下标
			if sum > target {
				third0 := third - 1
				// 计算的至一定要大 second
				for second < third0 && nums[third] == nums[third0] {
					third0--
				}
				third = third0
			}
			if sum < target {
				second0 := second + 1
				// 计算的至一定要小third
				for second0 < third && nums[second] == nums[second0] {
					second0++
				}
				second = second0
			}

		}
	}

	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// @lc code=end

