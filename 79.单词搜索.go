/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
var dirs = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	// 定义已访问的矩阵标记
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	// 每一个位置开始搜索
	for i, v := range board {
		for j := range v {
			// index记录已经找到的长度，而不是记录找到的词！！！
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false

}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, i, j int) bool {
	// 结束条件
	if index == len(word)-1 {
		return board[i][j] == word[index]
	}

	// 正常查找
	if board[i][j] == word[index] {
		visited[i][j] = true // 查找过
		// 四个方向开始寻找
		for _, dir := range dirs {
			// i j 不能变，多个循环都要用
			ni := i + dir[0]
			nj := j + dir[1]
			// 矩阵内 & **没有访问过** & 继续搜索
			if isInBoard(board, ni, nj) && !visited[ni][nj] && searchWord(board, visited, word, index+1, ni, nj) {
				return true
			}
		}

		// for i := 0; i < 4; i++ {
		// 	ni := i + dirs[i][0]
		// 	nj := j + dirs[i][1]
		// 	if isInBoard(board, ni, nj) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
		// 		return true
		// 	}
		// }

		visited[i][j] = false // 切记要去掉
	}
	return false
}

// @lc code=end
