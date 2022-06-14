from typing import *
from xml.dom import minidom
#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

# @lc code=start
class Solution:
    # 1、LCP(LCP(LCP(S1,S2),S3),…Sn)
    # def longestCommonPrefix(self, strs: List[str]) -> str:
    #     if not strs:
    #         return ""

    #     prefix, length = strs[0], len(strs)
    #     if length == 1:
    #         return prefix

    #     for i in range(1, length):
    #         prefix = self.lcp(prefix, strs[i])
    #         if prefix == "":
    #             break
    #     return prefix

    # def lcp(self, str1: str, str2: str) -> str:
    #     length, index = min(len(str1), len(str2)), 0
    #     while index < length and str1[index] == str2[index]:
    #         index += 1
    #     return str1[:index]

    # 2、竖向扫描，粗暴方法
    # def longestCommonPrefix(self, strs: List[str]) -> str:
    #     if not strs:
    #         return ""

    #     strs_len = len(strs)
    #     # if strs_len == 1:
    #     #     return strs[0]
    #     str_min_len = min([len(i) for i in strs])

    #     for i in range(str_min_len):

    #         for j in range(1, strs_len):
    #             if strs[0][i] != strs[j][i]:
    #                 return strs[0][0:i]
    #     return strs[0][:str_min_len]

    # 3、二分查找方式，递归查找，分而治之。
    def longestCommonPrefix(self, strs: List[str]) -> str:
        def lcp(strs1: List[str], start: int, end: int) -> str:
            if start == end:
                return strs1[start]
            
            middle = (start + end) // 2
            left_prefix, right_prefix = lcp(strs1, start, middle), lcp(strs1, middle + 1, end)
            min_length = min(len(left_prefix), len(right_prefix))
            for i in range(min_length):
                if left_prefix[i] != right_prefix[i]:
                    return left_prefix[:i]
            return left_prefix[:min_length]

        return "" if not strs else lcp(strs, 0, len(strs) - 1)
        
# @lc code=end

s = Solution()
print(s.longestCommonPrefix(["flower","flow","flight"]))
print(s.longestCommonPrefix(["ab", "a"]))