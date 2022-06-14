#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = sum_listnode = ListNode()
        carry  = 0
        while l1 or l2 or carry:
            val = carry
            if l1: l1, val = l1.next, val + l1.val
            if l2: l2, val = l2.next, val + l2.val

            carry, sum = divmod(val, 10)

            sum_listnode.next = ListNode(sum)
            sum_listnode = sum_listnode.next

        return head.next
# @lc code=end

