/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package main

import "fmt"

// @lc code=start

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if nums == nil || numsLen < 1 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	pre, maxRes := 0, nums[0]
	for i := 0; i < numsLen; i++ {
		pre = max(pre+nums[i], nums[i])
		maxRes = max(maxRes, pre)
	}
	return maxRes
}

// @lc code=end

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
