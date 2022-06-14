#
# @lc app=leetcode.cn id=5 lang=python3
#
# [5] 最长回文子串
#

# @lc code=start
class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        if n == 1:
            return s
        
        start = 0
        # 最少也是1个字母！！
        end = 0
        # dp建立二维数组，[i][j] 标识s[i:j+1]是否为回文
        dp = [[False] * n for _ in range(n)]
        # 本身是回文
        for i in range(n):
            dp[i][i] = True

        # 子串的大小，开始往外扩散
        for sub_s_L in range(2, n + 1):
            for l in range(n):
                # r - l + 1 = sub_s_L
                r = sub_s_L + l - 1
                if r >= n:
                    break
                    
                # 因为是从最小开始，所以开始判断两边
                if s[l] != s[r]:
                    dp[l][r] = False
                else:
                    # 小于三，可以直接判定dp表
                    if r - l < 3:
                        dp[l][r] = True
                    else:
                        # 直接根据状态转移公式计算
                        dp[l][r] = dp[l+1][r-1]
                

                if dp[l][r] and r - l > end - start:
                    start = l
                    end  = r
        return s[start: end+1]

# @lc code=end

# class Solution:
#     def longestPalindrome(self, s: str) -> str:
#         n = len(s)
#         start = end = 0
#         for i in range(n):
#             left1, right1 = self.extentLongPa(s, i, i)
#             left2, right2 = self.extentLongPa(s, i, i + 1)

#             len1 = right1 - left1
#             len2 = right2 - left2
#             len_long = end - start
#             if  len1 > len2 and len1 > len_long:
#                 start, end = left1, right1
#             elif len1 < len2 and len2 > len_long:
#                 start, end = left2, right2
            
#         return s[start:end+1]
    
#     def extentLongPa(self, s: str, l: int, r: int):
#         # 先对传入的参数进行校验
#         while l >=0 and r < len(s) and s[l] == s[r]:
#             l -= 1
#             r += 1
#         return l + 1, r - 1
