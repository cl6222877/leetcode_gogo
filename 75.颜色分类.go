/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {

	// 两次遍历
	// var index int
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		nums[index], nums[i] = nums[i], nums[index]
	// 		index++
	// 	}
	// }
	// for i := index; i < len(nums); i++ {
	// 	if nums[i] == 1 {
	// 		nums[index], nums[i] = nums[i], nums[index]
	// 		index++
	// 	}
	// }

	// 双指针，p0直向前，p2指向后
	length := len(nums)
	var p0, p2 int = 0, length - 1
	// 一个往前扫，一个往后i
	for i := 0; i <= p2; i++ {
		// 从后面开始，已经替换了所以不会死循环
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		// 从前面开始
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}

	}
}

// @lc code=end

// 题目末尾的 Follow up 提出了一个更高的要求，能否用一次循环解决问题？这题由于数字只会出现 0，1，2 这三个数字，所以用游标移动来控制顺序也是可以的。具体做法：0 是排在最前面的，所以只要添加一个 0，就需要放置 1 和 2。1 排在 2 前面，所以添加 1 的时候也需要放置 2 。至于最后的 2，只用移动游标即可。

// 这道题可以用计数排序，适合待排序数字很少的题目。用一个 3 个容量的数组分别计数，记录 0，1，2 出现的个数。然后再根据个数排列 0，1，2 即可。时间复杂度 O(n)，空间复杂度 O(K)。这一题 K = 3。

// 这道题也可以用一次三路快排。数组分为 3 部分，第一个部分都是 0，中间部分都是 1，最后部分都是 2 。
for i := 0; i <= p2; i++ {
        for ; i <= p2 && nums[i] == 2; p2-- {
            nums[i], nums[p2] = nums[p2], nums[i]
        }
        if nums[i] == 0 {
            nums[i], nums[p0] = nums[p0], nums[i]
            p0++
        }
}