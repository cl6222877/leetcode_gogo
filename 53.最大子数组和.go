/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	// 状态转义方程。f(i)=max{f(i−1)+nums[i],nums[i]}
	dp, ans := make([]int, length), nums[0]
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		ans = max(dp[i], ans)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end

func maxSubArray(nums []int) int {
	// 空间利用率O(1),不过会破坏nums的内容
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
