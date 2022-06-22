/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
// 这一题见lubladong的解法，容易理解
var note []int

func wordBreak(s string, wordDict []string) bool {
	note = make([]int, len(s))
	return trace(s, 0, wordDict)
}

func trace(s string, index int, wordDict []string) bool {
	if index == len(s) {
		return true
	}

	// 防止冗余计算
	if note[index] == 1 {
		return true
	}
	if note[index] == -1 {
		return false
	}

	// 遍历所有数，尝试匹配前缀
	for _, word := range wordDict {
		length := len(word)
		if index+length > len(s) {
			continue
		}
		// 不相等
		if s[index:index+length] != word {
			continue
		}
		// 相等，
		if trace(s, index+length, wordDict) {
			note[index] = 1
			return true
		}
	}
	note[index] = -1
	return false

}

// @lc code=end
