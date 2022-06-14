from curses.ascii import SO
from typing import *
#
# @lc app=leetcode.cn id=438 lang=python3
#
# [438] 找到字符串中所有字母异位词
#

# @lc code=start
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        # 根据diff处理
        s_len, p_len = len(s), len(p)
        ans = []
        if s_len < p_len:
            return ans

        diff_list = [0] * 26
        for i in range(p_len):
            diff_list[ord(p[i]) - 97] += 1
            diff_list[ord(s[i]) - 97] -= 1

        diff = [x !=0 for x in diff_list].count(True)
        if diff == 0:
            ans.append(0)

        # 注意，一下处理得出的结果实际是i+i的逻辑，所以最大不能超过相差，不能+1
        for i in range(s_len - p_len):
            # 去除i字母 diff的变化
            if diff_list[ord(s[i]) - 97] == -1:
                diff -= 1
            elif diff_list[ord(s[i]) - 97] == 0:
                diff += 1
            # 一定会增加1，之前是加入-1
            diff_list[ord(s[i]) - 97] += 1

            # 加上i+p_leng字母 diff的变化
            if diff_list[ord(s[i + p_len]) - 97] == 1:
                diff -= 1
            elif diff_list[ord(s[i + p_len]) - 97] == 0:
                diff += 1
                
            # 一定会减少一个,之前去除+1
            diff_list[ord(s[i + p_len]) - 97] -= 1

            if diff == 0:
                ans.append(i + 1)

        return ans
        
# @lc code=end

# 滑动窗口处理
# class Solution:
#     def findAnagrams(self, s: str, p: str) -> List[int]:
#         s_len, p_len = len(s), len(p)
#         ans = []

#         if s_len < p_len:
#             return ans

#         # 将字母装换成下标的形式存储个数
#         s_chars = [0] * 26
#         p_chars = [0] * 26
#         for i in range(p_len):
#             p_chars[ord(p[i]) - 97] += 1
#             s_chars[ord(s[i]) - 97] += 1

#         # 需要注意越界
#         # for i in range(s_len - p_len + 1):
#         #     if p_chars == s_chars:
#         #         ans.append(i)

#         #     if i + p_len + 1 > s_len: break
#         #     s_chars[ord(s[i]) - 97] -= 1
#         #     s_chars[ord(s[i + p_len]) - 97] += 1

#         if p_chars == s_chars:
#             ans.append(0)

#         for i in range(s_len - p_len):
#                 s_chars[ord(s[i]) - 97] -= 1
#                 s_chars[ord(s[i + p_len]) - 97] += 1
#                 if p_chars == s_chars:
#                     ans.append(i + 1)

#         return ans

s = Solution()
print(s.findAnagrams("baa", "aa"))