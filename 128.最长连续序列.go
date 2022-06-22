/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	// 模拟DP方法， 后半段更新数据不好理解
	res, numsMap := 0, map[int]int{} // num 对应的最长长度
	for _, num := range nums {
		// 如果此数还未进行处理，左边 或 右边都是没有对此数进行计算的，可以进行更新处理。
		if numsMap[num] == 0 {
			// 左边最长，右边最长
			leftSum, rightSum, sum := 0, 0, 0
			if numsMap[num-1] > 0 {
				leftSum = numsMap[num-1]
			}
			if numsMap[num+1] > 0 {
				rightSum = numsMap[num+1]
			}

			sum = leftSum + rightSum + 1 // 1 即此数本身
			// numsMap[num] = sum
			res = max(res, sum)

			// 更新数据最左边和最右边的长度数据（中间数据不会被用到了，不过这种节约了很多时间）,这种不好理解，我们直接全部更新了吧
			for left, right := num-leftSum, num+rightSum; left <= right; left++ {
				numsMap[left] = sum
			}

		} else {
			continue
		}

	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

// 暴力解决解决，代码见解法一。思路是把每个数都存在 map 中，先删去 map 中没有前一个数 nums[i]-1 也没有后一个数 nums[i]+1 的数 nums[i]，这种数前后都不连续。然后在 map 中找到前一个数 nums[i]-1 不存在，但是后一个数 nums[i]+1 存在的数，这种数是连续序列的起点，那么不断的往后搜，直到序列“断”了。最后输出最长序列的长度。
// 解法二，针对每一个 map 中不存在的数 n，插入进去都做 2 件事情。第一件事，先查看 n - 1 和 n + 1 是否都存在于 map 中，如果都存在，代表存在连续的序列，那么就更新 left，right 边界。那么 n 对应的这个小的子连续序列长度为 sum = left + right + 1。第二件事就是更新 left 和 right 左右边界对应的 length = sum。

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 暴力解法核心：前后没数字，一定不连续，可以去掉；前后有数据，中间值；前无数据，后有数据，开头值，从这里入手遍历。
	numsMap := map[int]bool{}
	// 初始化
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = true
	}

	for index := range numsMap {
		if !numsMap[index-1] && !numsMap[index+1] {
			delete(numsMap, index)
		}
	}

	// 如果全部都不符合，那最长就是1
	if len(numsMap) == 0 {
		return 1
	}
	res := 1
	for index := range numsMap {
		// 开始从起点遍历
		if !numsMap[index-1] && numsMap[index] {
			length := 1 // 起点有一个了
			tmp := index + 1
			// 下一个点存在
			for numsMap[tmp] {
				length++
				tmp++
			}
			res = max(res, length)
		}
	}
	return res
}
