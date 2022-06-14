from typing import *
#
# @lc app=leetcode.cn id=16 lang=python3
#
# [16] 最接近的三数之和
#

# @lc code=start

class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        length = len(nums)
        nums.sort()
        diff = 7**10

        # n**3 耗时太长。go用双指针方法
        for first in range(length):
            if first > 0 and nums[first] == nums[first - 1]:
                continue

            for second in range(first + 1, length):
                if second > first + 1 and nums[second] == nums[second - 1]:
                    continue
                # 这里因为abs函数不是单调的，不能用此方法，要不然会漏掉第一个小于diff后的 third
                # while second < third and abs(nums[first] + nums[second] + nums[third] - target) > diff:
                #     third -= 1
                # if second == third:
                #     break
                # if abs(nums[first] + nums[second] + nums[third] - target) < diff:
                #     diff = abs(nums[first] + nums[second] + nums[third] - target)
                #     best = nums[first] + nums[second] + nums[third]
                for third in range(second + 1, length):
                    if third > third + 1 and nums[third] == nums[third - 1]:
                        continue
                    if abs(nums[first] + nums[second] + nums[third] - target) < diff:
                        diff = abs(nums[first] + nums[second] + nums[third] - target)
                        best = nums[first] + nums[second] + nums[third] 
        
        return best



# @lc code=end

s = Solution()
print(s.threeSumClosest([1,1,-1,-1,3], -1))