/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
// diff
func findAnagrams(s string, p string) []int {
	sLength, pLength := len(s), len(p)
	ans := []int{}

	if sLength < pLength {
		return ans
	}

	// sChars, pChars := [26]int{}  错误写法，可以只定义不初始化，并且得有两个value对应
	var diffChars [26]int

	for i, val := range p {
		diffChars[s[i]-'a']--
		diffChars[val-'a']++
	}

	diff := 0
	for _, c := range diffChars {
		if c != 0 {
			diff++
		}
	}

	if diff == 0 {
		ans = append(ans, 0)
	}

	for i, val := range s[:sLength-pLength] {
		if diffChars[val-'a'] == -1 {
			diff--
		} else if diffChars[val-'a'] == 0 {
			diff++
		}
		diffChars[val-'a']++

		if diffChars[s[i+pLength]-'a'] == 1 {
			diff--
		} else if diffChars[s[i+pLength]-'a'] == 0 {
			diff++
		}
		diffChars[s[i+pLength]-'a']--

		if diff == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans

}

// @lc code=end
// func findAnagrams(s string, p string) []int {
// 	sLength, pLength := len(s), len(p)
// 	ans := []int{}

// 	if sLength < pLength {
// 		return ans
// 	}

// 	// sChars, pChars := [26]int{}  错误写法，可以只定义不初始化，并且得有两个value对应
// 	var sChars, pChars [26]int

// 	for i, val := range p {
// 		sChars[s[i]-'a']++
// 		pChars[val-'a']++
// 	}

// 	if sChars == pChars {
// 		ans = append(ans, 0)
// 	}

// 	// for i, _ := range sLength - pLength { 错误写法，不可遍历int，直接遍历差值
// 	for i, val := range s[:sLength-pLength] {
// 		// 去掉i处字母，加上i+plength处字母
// 		sChars[s[i]-'a']--
// 		sChars[s[i+pLength]]++

// 		if sChars == pChars {
// 			ans = append(ans, i+1)
// 		}
// 	}
// 	return ans

// }