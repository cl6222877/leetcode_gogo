/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 左右开始躺平
	flatten(root.Left)
	flatten(root.Right)

	// 左右开始拼接
	// 左边比较小
	right := root.Right
	root.Right = root.Left
	root.Left = nil

	// 找到最后一个节点
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right

}

// @lc code=end

