/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	// 初学者做这题很容易有误区：BST 不是左小右大么，那我只要检查 root.val > root.left.val 且 root.val < root.right.val 不就行了？
	// 这样是不对的，因为 BST 左小右大的特性是指 root.val 要比左子树的所有节点都更大，要比右子树的所有节点都小，你只检查左右两个子节点当然是不够的。
	return isValidbst(root, math.Inf(-1), math.Inf(1))

}

func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}

// @lc code=end

// 解法二，把 BST 按照左中右的顺序输出到数组中，如果是 BST，则数组中的数字是从小到大有序的，如果出现逆序就不是 BST
func isValidBST1(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
