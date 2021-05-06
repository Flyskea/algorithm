/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package main

import "fmt"

// @lc code=start
/*
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target, numsLen := sum/2, len(nums)
	dp := make([][]bool, numsLen+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for i := 0; i < numsLen; i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j] || dp[i][j-nums[i]]
			}
		}
	}
	return dp[numsLen][target]
} */

// 空间优化
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target, numsLen := sum/2, len(nums)
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < numsLen; i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[target] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}

// @lc code=end

func main() {
	fmt.Println(canPartition([]int{1, 2, 3}))
}
