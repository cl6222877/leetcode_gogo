/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	p, res := []int{}, [][]int{}
	used := make([]bool, len(nums))
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)

		// 不能直接用，可能内存中会进行修改
		// *res = append(*res, p)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			p = append(p, nums[i])
			(*used)[i] = true
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}

	}
	return
}

// @lc code=end
