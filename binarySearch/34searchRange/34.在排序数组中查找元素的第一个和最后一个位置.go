/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
package main

import "fmt"

// @lc code=start
func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []int{-1, -1}
	}
	low, high := lowSearch(nums, target), highSearch(nums, target)-1
	if low == numsLen || nums[low] != target {
		return []int{-1, -1}
	}
	return []int{low, high}
}

func lowSearch(nums []int, target int) int {
	low, high := 0, len(nums)
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func highSearch(nums []int, target int) int {
	low, high := 0, len(nums)
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low

}

// @lc code=end
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
