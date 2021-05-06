/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	cookie := 0
	childLen, cookieLen := len(g), len(s)
	for child < childLen && cookie < cookieLen {
		if g[child] <= s[cookie] {
			child++
		}
		cookie++
	}
	return child
}

// @lc code=end

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 2, 3}))
}
