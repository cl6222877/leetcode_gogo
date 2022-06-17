/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// 1.BFS 借助FIFO
	// if root == nil {
	// 	return [][]int{}
	// }

	// queue := []*TreeNode{root}
	// res := make([][]int, 0)

	// for len(queue) > 0 {
	// 	// 记住此时该层有多少个节点
	// 	l := len(queue)
	// 	tmp := make([]int, 0, l)
	// 	// 加入值，并且加入下一层数据
	// 	for i := 0; i < l; i++ {
	// 		tmp = append(tmp, queue[i].Val)
	// 		if queue[i].Left != nil {
	// 			queue = append(queue, queue[i].Left)
	// 		}
	// 		if queue[i].Right != nil {
	// 			queue = append(queue, queue[i].Right)
	// 		}

	// 	}
	// 	// 去掉已遍历节点
	// 	queue = queue[l:]
	// 	res = append(res, tmp)
	// }
	// return res

	// 2.DFS  记录节点&节点在那一层
	var res [][]int
	var dfsLevel func(root *TreeNode, level int)
	dfsLevel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 新增一层 res 1 2 3 …… level 0 1 2 ……
		if len(res) == level {
			res = append(res, []int{root.Val})
		} else {
			res[level] = append(res[level], root.Val)
		}
		dfsLevel(root.Left, level+1)
		dfsLevel(root.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}

// @lc code=end
