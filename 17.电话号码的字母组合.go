/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
// 这里可以理解为一棵树，回溯算法其实就是可以利用之前结果的一种是算法
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backTrack(digits, 0, "")
	return combinations
}

func backTrack(digits string, start int, combination string) {
	if start == len(digits) {
		combinations = append(combinations, combination)
	} else {
		chars := phoneMap[string(digits[start])]
		for i := 0; i < len(chars); i++ {
			backTrack(digits, start+1, combination+string(chars[i]))
		}
	}
}

// @lc code=end

