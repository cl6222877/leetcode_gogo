/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var first *ListNode = &ListNode{0, head}
	low, high := first, first

	// 先走块的节点,设置好间隔,注意，要找的是前n+1个节点
	for i := 0; i < n; i++ {
		high = high.Next
	}

	for ; high.Next != nil; high = high.Next {
		low = low.Next
	}
	low.Next = low.Next.Next
	return first.Next
}

// @lc code=end

