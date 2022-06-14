/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		curr := &strings.Builder{}
		// 计算前一次pre的数据
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			// 个数+字符
			curr.WriteString(strconv.Itoa(j - start))
			curr.WriteByte(prev[start])
		}
		prev = curr.String()
	}
	return prev
}

// @lc code=end

