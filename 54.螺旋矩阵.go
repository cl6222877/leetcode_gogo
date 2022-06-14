/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 {
		return ans
	}

	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	total := len(matrix[0]) * len(matrix)
	for total >= 1 {
		// →
		for i := left; i <= right && total >= 1; i++ {
			ans = append(ans, matrix[top][i])
			total--
		}
		top++
		// ↓
		for i := top; i <= bottom && total >= 1; i++ {
			ans = append(ans, matrix[i][right])
			total--
		}
		right--
		// ←
		for i := right; i >= left && total >= 1; i-- {
			ans = append(ans, matrix[bottom][i])
			total--
		}
		bottom--
		// ↑
		for i := bottom; i >= top && total >= 1; i-- {
			ans = append(ans, matrix[i][left])
			total--
		}
		left++
	}
	return ans
}

// @lc code=end
