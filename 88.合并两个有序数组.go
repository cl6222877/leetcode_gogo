/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// i 是 数组中的下标,从后往前一个一个填充， m n 可以理解为双指针
	for i := m + n; m > 0 && n > 0; i-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[i-1] = nums2[n-1]
			n--
		} else {
			nums1[i-1] = nums1[m-1]
			m--
		}
	}
	// 移动了多少次n说明nums1中有多少次的值无用了，num2还有剩余直接填充
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// @lc code=end

