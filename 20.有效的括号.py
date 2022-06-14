#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    # 插入左括号，判断右括号
    def isValid(self, s: str) -> bool:
        sLen = len(s)
        if sLen % 2 !=0: return False

        stack = list()
        rightMap = {
            ')': '(',
		    '}': '{',
		    ']': '[',
        }

        for i in s:
            if i in rightMap:
                if len(stack) == 0 or rightMap[i] != stack[-1]:
                    return False
                stack = stack[:-1]
            else:
                stack.append(i)
        return len(stack) == 0



# @lc code=end

