/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */
package main

import "fmt"

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	numsLen := len(nums)
	if numsLen < 3 {
		return 0
	}
	dp := make([]int, numsLen)
	for i := 2; i < numsLen; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	res := 0
	// 任意等差数列连在一起也是等差数列
	for _, v := range dp {
		res += v
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
