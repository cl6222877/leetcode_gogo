/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环. 
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	l, r := 0, length - 1
	for l <= r {
		var mid int = (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		// 在数据大的区间内
		if nums[mid] >= nums[0] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		// 在数据小的区间内
		} else {
			if nums[mid] < target && target <= nums[length - 1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
// @lc code=end


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target:
                return mid
            if nums[0] <= nums[mid]:
                if nums[0] <= target < nums[mid]:
                    r = mid - 1
                else:
                    l = mid + 1
            else:
                if nums[mid] < target <= nums[len(nums) - 1]:
                    l = mid + 1
                else:
                    r = mid - 1
        return -1
