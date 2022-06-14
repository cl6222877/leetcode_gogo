from typing import *
#
# @lc app=leetcode.cn id=11 lang=python3
#
# [11] 盛最多水的容器
#

# @lc code=start
class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, right = 0, len(height) - 1
        ans = 0
        while left < right:
            left_num = height[left]
            right_num = height[right]
            length = right - left
            if left_num > right_num:
                min_num = right_num
                right -= 1
            else:
                min_num = left_num
                left += 1
            area = min_num * length
            ans = max(ans ,area)
        return ans
# @lc code=end
