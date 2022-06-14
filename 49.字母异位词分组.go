/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	mp := map[[26]int][]string{}
	for _, str := range strs {
		cur := [26]int{}
		for _, ch := range str {
			cur[ch-'a']++
		}

		mp[cur] = append(mp[cur], str)
	}

	// 铺开
	var ans [][]string
	// ans := make([][]string, 0, len(mp))
	for _, substrs := range mp {
		ans = append(ans, substrs)
	}

	return ans
}

// @lc code=end

