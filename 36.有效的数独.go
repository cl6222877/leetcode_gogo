/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int

	for i, row := range board {
		for j, val := range row {
			if val == '.' {
				continue
			}
			valNum := val - '1'
			rows[i][valNum]++
			columns[j][valNum]++
			subboxes[i/3][j/3][valNum]++

			if rows[i][valNum] > 1 || columns[j][valNum] > 1 || subboxes[i/3][j/3][valNum] > 1 {
				return false
			}

		}
	}
	return true
}

// @lc code=end

