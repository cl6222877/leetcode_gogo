/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
// 1、横向扫描
// func longestCommonPrefix(strs []string) string {
// 	length := len(strs)
// 	if length == 0 {
// 		return ""
// 	}
// 	prefix := strs[0]
// 	for i := 1; i < length; i++ {
// 		prefix = lcp(prefix, strs[i])
// 		if len(prefix) == 0 {
// 			break
// 		}
// 	}
// 	return prefix

// }

// func lcp(strs1 string, strs2 string) string {
// 	length := len(strs1)
// 	for i := 0; i < length; i++ {
// 		if i == len(strs2) || strs1[i] != strs2[i] {
// 			fmt.Println(i)
// 			return strs1[:i]
// 		}
// 	}
// 	// 如果为“” 或者全部相等
// 	return strs1
// }

// 2、纵向暴力扫描,其实就是所有的进行对比，类似方法1的lcp
// func longestCommonPrefix(strs []string) string {
// 	length := len(strs)
// 	if length == 0 {
// 		return ""
// 	}
// 	for i := 0; i < len(strs[0]); i++ {
// 		for j := 1; j < length; j++ {
// 			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
// 				return strs[0][:i]
// 			}
// 		}
// 	}
// 	//  “” 或者全部相等
// 	return strs[0]

// }

// 3、分而治之
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	var lcp func(int, int) string
	lcp = func(start, end int) string {
		// 1 结束条件
		if start == end {
			return strs[start]
		}

		// 2 迭代转移方程
		middle := (start + end) / 2
		left, right := lcp(start, middle), lcp(middle+1, end)
		for i := 0; i < len(left); i++ {
			if i == len(right) || left[i] != right[i] {
				return left[:i]
			}
		}
		// 如果为“” 或者全部相等
		return left

	}

	return lcp(0, len(strs)-1)

}

// @lc code=end

