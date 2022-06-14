/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
type NewInts [][]int

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	return n[i][0] < n[j][0]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return intervals
	}

	// 排序,利用这种办法，估计源码是n^2
	sort.Sort(NewInts(intervals))

	ans := [][]int{}
	for i := 0; i < length; i++ {
		// 没有重合，新增
		if len(ans) == 0 || ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		} else {
			// 有重合，如果末尾比较大，更新之
			if ans[len(ans)-1][1] < intervals[i][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		}
	}

	return ans
}

// @lc code=end

