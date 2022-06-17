/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 1.BFS 用一个队列，
	// if root == nil {
	// 	return [][]int{}
	// }
	// var res [][]int
	// queue := []*TreeNode{root}
	// leftToRight := true
	// for len(queue) > 0 {
	// 	l := len(queue)
	// 	tmp := make([]int, 0)
	// 	if leftToRight {
	// 		for i := 0; i < l; i++ {
	// 			tmp = append(tmp, queue[i].Val)
	// 		}
	// 	} else {
	// 		for j := l - 1; j >= 0; j-- {
	// 			tmp = append(tmp, queue[j].Val)
	// 		}
	// 	}
	// 	// 这里要注意，加入节点的顺序不能变
	// 	for i := 0; i < l; i++ {
	// 		if queue[i].Left != nil {
	// 			queue = append(queue, queue[i].Left)
	// 		}
	// 		if queue[i].Right != nil {
	// 			queue = append(queue, queue[i].Right)
	// 		}
	// 	}
	// 	leftToRight = !leftToRight
	// 	res = append(res, tmp)
	// 	queue = queue[l:]

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
		} else if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			// 数组后...打散，参数中...string 标识可以传递多个string变量
			res[level] = append([]int{root.Val}, res[level]...)
		}
		dfsLevel(root.Left, level+1)
		dfsLevel(root.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}

// @lc code=end
