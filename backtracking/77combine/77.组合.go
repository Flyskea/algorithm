/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
package main

import "fmt"

// @lc code=start
func backtracking(res *[][]int, comb *[]int, pos, n, k int) {
	if len(*comb) == k {
		tmp := []int{}
		tmp = append(tmp, *comb...)
		*res = append(*res, tmp)
		return
	}
	for i := pos; i <= n; i++ {
		*comb = append(*comb, i)
		backtracking(res, comb, i+1, n, k)
		*comb = (*comb)[:len(*comb)-1]
	}
}

func combine(n int, k int) [][]int {
	var res [][]int
	var comb []int
	backtracking(&res, &comb, 1, n, k)
	return res
}

// @lc code=end

func main() {
	fmt.Println(combine(4, 2))
}
