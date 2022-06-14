/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// length := len(nums)
	// if length == 0 {
	// 	return []int{-1, -1}
	// }

	// lindex, rindex := 0, length-1
	// l, r := -1, -1
	// for lindex <= rindex {
	// 	mid := (lindex + rindex) / 2
	// 	if nums[mid] == target {
	// 		l, r = target, target
	// 		// 这里可能会变成O(n)
	// 		for start := mid - 1; start >= 0; start-- {
	// 			if nums[start] == target {
	// 				l = start
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		for end := mid + 1; end <= length-1; end++ {
	// 			if nums[end] == target {
	// 				l = end
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 	}

	// 	if nums[mid] > target {
	// 		rindex = mid - 1
	// 	}

	// 	if nums[mid] < target {
	// 		lindex = mid + 1
	// 	}
	// }
	// return []int{l, r}

	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			// 这里是关键，一定要往左边继续找
			r = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			// 这里是关键，一定要往右边继续找
			l = mid + 1
		}
	}
	return -1
}

// @lc code=end

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] < target {
			l = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			// 这里是关键，一定要往右边继续找
			r = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 2)
		if nums[mid] > target {
			r = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			// 这里是关键，一定要往右边继续找
			l = mid + 1
		}
	}
	return -1
}
