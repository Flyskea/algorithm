/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package main

import "fmt"

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum || (target+sum)%2 == 1 {
		return 0
	}
	tar := (sum + target) / 2
	dp := make([]int, tar+1)
	dp[0] = 1
	for _, v := range nums {
		for i := tar; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[tar]
}

// @lc code=end

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
