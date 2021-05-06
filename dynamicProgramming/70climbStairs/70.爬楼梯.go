/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package main

import "fmt"

// @lc code=start
/* func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	res := dp[n-1]
	return res
} */

// 空间优化

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre1 := 2
	pre2 := 1
	cur := 0
	for i := 2; i < n; i++ {
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

// @lc code=end

func main() {
	fmt.Println(climbStairs(4))
}
