#
# @lc app=leetcode.cn id=8 lang=python3
#
# [8] 字符串转换整数 (atoi)
#

# @lc code=start
from enum import auto


INT_MAX = 2**31 - 1
INT_MIN = -2**31

class Automaton:
    def __init__(self) -> None:
        self.state = 'start'
        self.sign = 1
        self.ans = 0
        self.table = {
            # 初始状态 -> 空格，符号+/-，数字，other
            'start': ['start', 'signed', 'in_number', 'end'],
            'signed': ['end', 'end', 'in_number', 'end'],
            'in_number': ['end', 'end', 'in_number', 'end'],
            'end': ['end', 'end', 'end', 'end'],
        }

    def get_col(self, c: chr) -> int:
        """
        根据字符类型，判断下一步
        """
        if c.isspace():
            return 0
        elif c == '+' or c == '-':
            return 1
        elif c.isdigit():
            return 2
        else:
            return 3
        
    def get(self, c: chr) -> None:
        self.state = self.table[self.state][self.get_col(c)]
        if self.state == 'signed':
            self.sign = 1 if c == '+' else -1

        if self.state == 'in_number':
            self.ans = self.ans * 10 + int(c)
            # 注意，这里的self.ans是无符号的
            self.ans = min(self.ans, INT_MAX) if self.sign == 1 else min(self.ans, -INT_MIN)


class Solution:
    def myAtoi(self, s: str) -> int:
        auto = Automaton()
        for ch in s:
            # 如果已经end了，没有必要继续走了
            if auto.state == 'end':
                break
            auto.get(ch)
        return auto.sign * auto.ans


# @lc code=end

