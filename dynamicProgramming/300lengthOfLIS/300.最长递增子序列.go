/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */
package main

import "fmt"

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)
	maxLen := 0
	for i := 0; i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen + 1
}

// @lc code=end

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
