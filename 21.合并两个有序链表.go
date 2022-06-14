/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**index1, index2 := 0, 0
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	// 回溯递归，可以使用前一次结果，继续调用方法
// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val < list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}
// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }

// 双指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 擦了，链表的长度不可估量！！

	var head *ListNode = &ListNode{Next: list1}
	tail := head
	p1, p2 := list1, list2

	// 一次只移动一个，<=可以归并一起，记住末尾的tail节点
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			tail.Next = p2
			p2 = p2.Next
		} else {
			tail.Next = p1
			p1 = p1.Next
		}
		tail = tail.Next
	}

	if p1 != nil {
		tail.Next = p1
	}

	if p2 != nil {
		tail.Next = p2
	}

	return head.Next

}

// @lc code=end

