/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry

		carry = sum / 10
		sum = sum % 10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 最后一次计算如果有的话
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{Val: 0}
// 	n1, n2, carry, current := 0, 0, 0, head
// 	for l1 != nil || l2 != nil || carry != 0 {
// 		if l1 == nil {
// 			n1 = 0
// 		} else {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 == nil {
// 			n2 = 0
// 		} else {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}
// 		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
// 		current = current.Next
// 		carry = (n1 + n2 + carry) / 10
// 	}
// 	return head.Next
// }
// func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
// 	var tail *ListNode
// 	carry := 0
// 	for l1 != nil || l2 != nil {
// 		n1, n2 := 0, 0
// 		if l1 != nil {
// 			n1 = l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			n2 = l2.Val
// 			l2 = l2.Next
// 		}
// 		sum := n1 + n2 + carry
// 		sum, carry = sum%10, sum/10
// 		if head == nil {
// 			head = &ListNode{Val: sum}
// 			tail = head
// 		} else {
// 			tail.Next = &ListNode{Val: sum}
// 			tail = tail.Next
// 		}
// 	}
// 	if carry > 0 {
// 		tail.Next = &ListNode{Val: carry}
// 	}
// 	return
// }

// @lc code=end
