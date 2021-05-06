/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */
package main

import "fmt"

// @lc code=start
func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	var mid int
	var numsLenIsEven bool
	for low < high {
		mid = low + (high-low)/2
		numsLenIsEven = (high-mid)%2 == 0
		if nums[mid+1] == nums[mid] {
			if numsLenIsEven {
				low = mid + 2
			} else {
				high = mid - 1
			}
		} else if nums[mid-1] == nums[mid] {
			if numsLenIsEven {
				high = mid - 2
			} else {
				low = mid + 1
			}
		} else {
			return nums[mid]
		}
	}
	return nums[low]
}

// @lc code=end

func singleNonDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {
	fmt.Println(singleNonDuplicate1([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
