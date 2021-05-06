/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */
package main

import "fmt"

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerLen := len(flowerbed)
	if n == 0 {
		return true
	}
	for i := 0; i < flowerLen; i += 2 {
		if flowerbed[i] == 0 {
			if i == flowerLen-1 || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
		if n == 0 {
			return true
		}
	}
	return n <= 0
}

// @lc code=end

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}
