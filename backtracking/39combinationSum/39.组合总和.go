/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	c := []int{}
	backtracking(candidates, target, 0, c, &res)
	return res
}

func backtracking(candidates []int, target, index int, total []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			*res = append(*res, append([]int{}, total...))
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		total = append(total, candidates[i])
		backtracking(candidates, target-candidates[i], i, total, res)
		total = total[:len(total)-1]
	}
}

// @lc code=end

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
