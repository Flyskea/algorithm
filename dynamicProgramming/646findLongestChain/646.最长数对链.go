/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

// @lc code=start
/* DP
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func findLongestChain(pairs [][]int) int {
	rows := len(pairs)
	dp := make([]int, rows+1)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1])
	})
	res := 0
	for i := 1; i < rows; i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res + 1
} */

// 贪心
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})
	res := 0
	cur := math.MinInt32
	for i := 0; i < len(pairs); i++ {
		if cur < pairs[i][0] {
			cur = pairs[i][1]
			res++
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(findLongestChain([][]int{{2, 3}, {3, 4}, {1, 2}}))
}
