from cgitb import reset
from turtle import towards
from typing import *
#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#

# @lc code=start
# https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        length = len(nums)
        
        if length < 3:
            return []
        nums.sort()
        res = []
        for i in range(length):
            if i > 0 and nums[i-1] == nums[i]:
                continue

            two_ress = self.twoSum(nums, i+1, 0 - nums[i])
            for two_res in two_ress:
                two_res.append(nums[i])
                res.append(two_res)
        return res 
        
    def twoSum(self, nums: List[int], start: int, target: int) -> List[List[int]]:
        length = len(nums)
        if length < 2:
            return []
        low, high = start, length - 1
        ans = []

        while low < high:
            left_num = nums[low]
            right_num = nums[high]
            sum = left_num + right_num
            if sum < target:
                while low < high and nums[low] == left_num:
                    low += 1
            elif sum > target:
                while low < high and nums[high] == right_num:
                    high -= 1 
            else:
                ans.append([nums[low], nums[high]])
                while low < high and nums[low] == left_num:
                    low += 1
                while low < high and nums[high] == right_num:
                    high -= 1 
        return ans
# @lc code=end

a = Solution()
print(a.threeSum([-1,0,1,2,-1,-4]))