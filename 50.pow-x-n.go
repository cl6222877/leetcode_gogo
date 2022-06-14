/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {

	if n > 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)

}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := myPow(x, n/2)
	if n%2 == 0 {
		return res * res
	}
	return res * res * x
}

// @lc code=end

