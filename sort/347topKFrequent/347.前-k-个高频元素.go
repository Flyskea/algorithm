/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	keymap := make(map[int]int)
	maxn := math.MinInt64
	for _, i := range nums {
		if _, ok := keymap[i]; !ok {
			keymap[i] = 1
		} else {
			keymap[i]++
		}
		if keymap[i] > maxn {
			maxn = keymap[i]
		}
	}
	hashtop := make([][]int, maxn+1)
	for k, v := range keymap {
		hashtop[v] = append(hashtop[v], k)
	}
	res := make([]int, 0)
	for i := maxn; i >= 0; i-- {
		res = append(res, hashtop[i]...)
		k -= len(hashtop[i])
		if k == 0 {
			break
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}
