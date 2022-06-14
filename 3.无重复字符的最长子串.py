#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        s_set = set()
        n = len(s)
        rk, ans = -1, 0
        for i in range(n):
            # 做指针右移动，去掉
            if i != 0:
                s_set.remove(s[i - 1])
            while rk + 1 < n and s[rk + 1] not in s_set:
                s_set.add(s[rk + 1])
                rk += 1

            ans = max(ans, rk - i + 1)
        return ans
# @lc code=end

