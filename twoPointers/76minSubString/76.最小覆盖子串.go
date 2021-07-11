/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
package main

import (
	"fmt"
)

// @lc code=start
func minWindow(s string, t string) string {
	chars := make([]int, 128)
	flag := make([]bool, 128)
	for i := range t {
		flag[t[i]] = true
		chars[t[i]]++
	}
	var cnt, l, minL, minSize int
	minSize = len(s) + 1
	for r := 0; r < len(s); r++ {
		if flag[s[r]] {
			chars[s[r]]--
			if chars[s[r]] >= 0 {
				cnt++
			}
			for cnt == len(t) {
				if r-l+1 < minSize {
					minL = l
					minSize = r - l + 1
				}
				if flag[s[l]] {
					chars[s[l]]++
					if chars[s[l]] > 0 {
						cnt--
					}
				}
				l++
			}
		}
	}
	if minSize > len(s) {
		return ""
	}
	return s[minL : minL+minSize]
}

// @lc code=end

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println((1 / 3) * 3)
}
