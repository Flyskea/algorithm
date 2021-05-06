/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */
package main

import "fmt"

// @lc code=start
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return validPalindromeHelper(s, i, j-1) || validPalindromeHelper(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func validPalindromeHelper(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

func main() {
	fmt.Println(validPalindrome("abc"))
}
