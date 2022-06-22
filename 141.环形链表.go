/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 快慢指针，同时环形一定相遇
	var fast, slow *ListNode = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 有可能第一次只有一个节点，这里就成功了理论上，要排除这种情况
		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=end

