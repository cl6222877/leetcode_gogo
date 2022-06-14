/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	mapp := make(map[int]int, n)
	var dp func(int) int
	dp = func(n int) int {
		if n <= 2 {
			return n
		}

		if mapp[n] > 0 {
			return mapp[n]
		}

		mapp[n] = dp(n-1) + dp(n-2)
		return mapp[n]
		// return dp(n-1) + dp(n-2)
	}
	return dp(n)

}

// @lc code=end

