/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowZero, colZero := false, false
	// 判断第一行第一列本身有没有0
	for _, v := range matrix[0] {
		if v == 0 {
			rowZero = true
			break
		}
	}
	for _, v := range matrix {
		if v[0] == 0 {
			colZero = true
			break
		}
	}

	// 根据除开第一行第一列的数值，更改第一行第一列的对应值为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	// 处理[1:]行全部置 0
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	// 处理[1:]列全部置 0
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据原先记录的标记，更改第一行第一列
	if rowZero {
		// for _, v := range matrix[0] {
		// 	v = 0
		// }
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if colZero {
		for _, v := range matrix {
			v[0] = 0
		}
	}

	// return matrix
}

// @lc code=end

// 此题考查对程序的控制能力，无算法思想。题目要求采用原地的算法，所有修改即在原二维数组上进行。在二维数组中有 2 个特殊位置
// 一个是第一行，一个是第一列。它们的特殊性在于，它们之间只要有一个 0，它们都会变为全 0 。
// 先用 2 个变量记录这一行和这一列中是否有 0，防止之后的修改覆盖了这 2 个地方。然后除去这一行和这一列以外的部分判断是否有 0，
// 如果有 0，将它们所在的行第一个元素标记为 0，所在列的第一个元素标记为 0 。最后通过标记，将对应的行列置 0 即可。
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}