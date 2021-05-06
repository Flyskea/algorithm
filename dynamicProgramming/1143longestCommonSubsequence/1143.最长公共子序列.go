/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
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

func longestCommonSubsequence(text1 string, text2 string) int {
	text1Len, text2Len := len(text1), len(text2)
	if text1Len <= 0 || text2Len <= 0 {
		return 0
	}
	dp := make([][]int, text1Len+1)
	for i := range dp {
		dp[i] = make([]int, text2Len+1)
	}
	for i := 1; i <= text1Len; i++ {
		for j := 1; j <= text2Len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[text1Len][text2Len]
}

// @lc code=end

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
