/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
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

func robCore(nums []int) int {
	numsLen := len(nums)
	if nums == nil || numsLen < 1 {
		return 0
	}
	dp := make([]int, numsLen+1)
	dp[1] = nums[0]
	for i := 2; i <= numsLen; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	res := dp[numsLen]
	return res
}
func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	return max(robCore(nums[1:]), robCore(nums[:len(nums)-1]))
}

// @lc code=end

func main() {
	a := []int{2, 3, 2}
	fmt.Println(a[:len(a)-1])
	fmt.Println(rob([]int{1}))
}
