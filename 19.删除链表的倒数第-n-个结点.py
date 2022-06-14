#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第 N 个结点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    # 第一种，放入栈中，pop去拿到n的前一位,但是这个空间复杂度为o~N
    # def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        # stack = list()
        # first = ListNode(0, head)
        # cur = first
        # while cur:
        #     stack.append(cur)
        #     cur = cur.next
        
        # for i in range(n):
        #     stack.pop()
        
        # prev = stack[-1]
        # prev.next = prev.next.next
        # return first.next

    # 第二种，快慢指针。先走n步，第二个指针开始走，一直到第一个指针走到最后一个。保持这两个指针为n
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        first = ListNode(0, head)
        low, high = first, first
        # 设置中间的间隔+1，因为要找到的是前一个节点
        for _ in range(n):
            high = high.next

        # while到最后一个节点
        while high.next:
            low = low.next
            high = high.next

        low.next = low.next.next
        return first.next



# @lc code=end

