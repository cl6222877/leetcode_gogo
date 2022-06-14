#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
#

# @lc code=start
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        if n == 0:
            return []
        ans = []
        # 注意这里用的是list来存储,也可以用string吧等下试试
        def backtrack(tmp: list, lindex: int, rindex: int):
            if len(tmp) == 2 * n:
                ans.append(''.join(tmp))
                return

            if lindex > 0:
                tmp.append('(')
                backtrack(tmp, lindex-1, rindex)
                # 因为这里修改了tmp结构，不影响这一次)的添加
                tmp.pop()
            if rindex > 0 and lindex < rindex:
                tmp.append(')')
                backtrack(tmp, lindex, rindex-1)
                tmp.pop()

        backtrack([], n, n)
        return ans
            

# @lc code=end

