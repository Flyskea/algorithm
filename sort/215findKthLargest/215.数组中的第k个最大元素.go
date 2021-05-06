/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

import "fmt"

// @lc code=start
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	p := partition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k < p {
		return selection(nums, l, p-1, k)
	} else {
		return selection(nums, p+1, r, k)
	}

}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

// @lc code=end

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
