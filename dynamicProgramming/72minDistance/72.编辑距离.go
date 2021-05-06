/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
package main

import "fmt"

// @lc code=start

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	wordLen1, wordLen2 := len(word1), len(word2)
	dp := make([][]int, wordLen1+1)
	for i := 0; i <= wordLen1; i++ {
		dp[i] = make([]int, wordLen2+1)
	}
	for i := 0; i <= wordLen1; i++ {
		for j := 0; j <= wordLen2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[wordLen1][wordLen2]
}

// @lc code=end

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
