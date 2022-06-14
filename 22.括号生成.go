/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	ans := []string{}

	backtrack(n, n, "", &ans)

	return ans
}

func backtrack(lindex int, rindex int, tmp string, ans *[]string) {
	if lindex == 0 && rindex == 0 {
		*ans = append(*ans, tmp)
		return
	}

	// 优先加(看看情况
	if lindex > 0 {
		backtrack(lindex-1, rindex, tmp+"(", ans)
	}

	// 左边符号一定多余右边符号,继续看加有括号的情况
	if rindex > 0 && lindex < rindex {
		backtrack(lindex, rindex-1, tmp+")", ans)
	}
}

// @lc code=end

