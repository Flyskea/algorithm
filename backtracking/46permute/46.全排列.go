/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package main

import "fmt"

// @lc code=start
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func backtracking(nums []int, level int, res *[][]int) {
	numsLen := len(nums)
	if level == numsLen-1 {
		tmp := make([]int, 0)
		tmp = append(tmp, nums...)
		*res = append(*res, tmp)
		return
	}
	for i := level; i < numsLen; i++ {
		swap(&nums[i], &nums[level])
		backtracking(nums, level+1, res)
		swap(&nums[i], &nums[level])
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	backtracking(nums, 0, &res)
	return res
}

// @lc code=end

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
