/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */
package main

import "fmt"

// @lc code=start
func twoSum(numbers []int, target int) []int {
	numsLen := len(numbers)
	var i, j int
	for i, j = 0, numsLen-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}

// @lc code=end
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
