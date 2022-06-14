/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	// 先对角
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	// 再水平翻转
	for i := 0; i < len(matrix); i++ {
		for left, right := 0, len(matrix)-1; left < right; {
			tmp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = tmp
			left++
			right--
		}
	}

}

// @lc code=end

