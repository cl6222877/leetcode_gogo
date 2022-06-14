/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	ROMAN_INT_MAP := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0
	for index, val := range s {
		num := ROMAN_INT_MAP[byte(val)]
		after_num := 0
		if index != len(s)-1 {
			after_num = ROMAN_INT_MAP[s[index+1]]
		}

		if num < after_num {
			ans -= num
		} else {
			ans += num
		}
	}

	return ans
}

// @lc code=end

