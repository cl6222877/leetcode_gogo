/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	// 这种写法会夺走好多层（左右层级不一致） 要正确找到结束的标记
	// if left == nil && right == nil {
	// 	return true
	// }

	if left == nil || right == nil {
		return right == left // 其实就是确定是否都会nil
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

// @lc code=end

// 迭代方法
func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}

// 中序遍历
func BFS(root *TreeNode, ans *[]int) {
	BFS(root.Left, ans)
	ans = append(ans, root.Val)
	BFS(root.Right, ans)
}

