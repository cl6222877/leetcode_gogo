#
# @lc app=leetcode.cn id=13 lang=python3
#
# [13] 罗马数字转整数
#

# @lc code=start
ROMAN_MAP_INT = {
    'I' : 1,
    'V' : 5,
    'X' : 10,
    'L' : 50,
    'C' : 100,
    'D' : 500,
    'M' : 1000,
}
class Solution:
    def romanToInt(self, s: str) -> int:
        ans = 0
        for i in range(len(s)):
            num =  ROMAN_MAP_INT[s[i]]
            after_num = 0 if i == len(s) - 1 else ROMAN_MAP_INT[s[i + 1]]

            if after_num > num:
                ans -= num
            else:
                ans += num
        return ans
# @lc code=end

