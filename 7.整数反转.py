
#
# @lc app=leetcode.cn id=7 lang=python3
#
# [7] 整数反转
#

# @lc code=start
from unicodedata import digit


class Solution:
    def reverse(self, x: int) -> int:
        INT_MIN = -2**31
        INT_MAX = 2**31 - 1

        rev = 0
        while x != 0:
            # int_min // 10 + 1 <= rev <= int_max // 10
            if rev < INT_MIN // 10 + 1 or rev > INT_MAX // 10:
                return 0
            
            # 个位上的数，*可能是0，不能随便-10
            digit = x % 10 # python3会特殊处理
            if x < 0 and digit != 0:
                digit = digit - 10
            
            # 剩余的数,可能刚好整除,不能随便+1
            x = x // 10 # python3会向下取整
            if x < 0 and digit !=0:
                x += 1

            # 翻转的数
            rev = rev * 10 + digit
        return rev

# @lc code=end

