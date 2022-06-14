/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	// 双指针办法，只能解决完全对称的符号，它有可能是这样子的"()[]{}"，所以此方法不适合,老老实实用栈的方式吧。
	// length := len(s)
	// if length%2 != 0 {
	// 	return false
	// }
	// var leftAll map[byte]byte = map[byte]byte{
	// 	'(': ')',
	// 	'{': '}',
	// 	'[': ']',
	// }
	// left, right := 0, length-1
	// for left < right {
	// 	// 这里需要建立映射值
	// 	if leftAll[s[left]] > 0 && leftAll[s[left]] == s[right] {
	// 		left++
	// 		right--
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true

	length := len(s)
	if length%2 != 0 {
		return false
	}
	var rightMap map[byte]byte = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte = []byte{}
	for i := 0; i < length; i++ {
		// 如果是右括号，并且栈中对应的一定是相反括号
		if rightMap[s[i]] > 0 {
			stackLen := len(stack)
			if stackLen == 0 || rightMap[s[i]] != stack[stackLen-1] {
				return false
			}
			stack = stack[:stackLen-1]

		} else {
			stack = append(stack, s[i])
		}
	}

	// 如果stack中还存在值说明不对应
	if len(stack) != 0 {
		return false
	}
	return true

}

// @lc code=end

