/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package main

import "fmt"

// @lc code=start
/* func rob(nums []int) int {
	numsLen := len(nums)
	if nums == nil || numsLen < 1 {
		return 0
	}
	dp := make([]int, numsLen+1)
	dp[1] = nums[0]
	for i := 2; i <= numsLen; i++ {
		if dp[i-1] > (nums[i-1] + dp[i-2]) {
			dp[i] = dp[i-1]
		} else {
			dp[i] = nums[i-1] + dp[i-2]
		}
	}
	return dp[numsLen]
} */

// 空间优化
func rob(nums []int) int {
	numsLen := len(nums)
	if nums == nil || numsLen < 1 {
		return 0
	}
	pre1, pre2, cur := 0, 0, 0
	for i := 0; i < numsLen; i++ {
		if pre1 > (nums[i] + pre2) {
			cur = pre1
		} else {
			cur = nums[i] + pre2
		}
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

// @lc code=end

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
