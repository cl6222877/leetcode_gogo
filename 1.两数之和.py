from typing import *
#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = {}
        for i, x in enumerate(nums):
            less = target - x

            if less in map.keys():
                return [i, map[less]]
            map[x] = i

# @lc code=end

test = Solution()
print(test.twoSum([123,321,32], 155))