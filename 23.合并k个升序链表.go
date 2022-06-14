/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func mergeKLists(lists []*ListNode) *ListNode {
// 	// 顺序
// 	var ans *ListNode
// 	for i := 0; i < len(lists); i++ {
// 		ans = mergeTwoLists2(ans, lists[i])
// 	}
// 	return ans
// }

// 回溯
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists2(merge(lists, l, mid), merge(lists, mid+1, r))
}

// 回溯
func mergeTwoLists(alist *ListNode, blist *ListNode) *ListNode {
	if alist == nil {
		return blist
	}

	if blist == nil {
		return alist
	}

	if alist.Val < blist.Val {
		alist.Next = mergeTwoLists(alist.Next, blist)
		return alist
	} else {
		blist.Next = mergeTwoLists(alist, blist.Next)
		return blist
	}

}

// 双指针
func mergeTwoLists2(alist *ListNode, blist *ListNode) *ListNode {
	var head *ListNode = &ListNode{0, alist}
	tail := head
	// 不破坏结构
	p1, p2 := alist, blist

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

	// 末尾加入后面多出的一串值
	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}

	return head.Next
}

// @lc code=end

