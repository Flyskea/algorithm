/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
package main

import (
	"fmt"
)

// 原数组中最多存在一处 nums[i] > nums[i+1]的情况。举个反例：
// 假设存在 ii，使得 nums[i] > nums[i+1] > nums[i+2]nums[i]>nums[i+1]>nums[i+2]，在这个不等式中，
// 存在两处前面元素大于后面元素的情况，我们可以发现，无论修改哪个元素，都无法使得上面不等式变为非递减的关系。

// 也就说，我们可以通过判断数组中存在 nums[i] > nums[i+1]的情况有多少处，
// 进而来判断是否能够最多改变 1 个元素，使得数组变为非递减数列。
// 考虑到前面遍历后都满足我们条件，即非递减，那么只要去和上一个，即 ums[i-1]去比较， 它满足 nums[i-1] <= nums[i] <= nums[i+1]
// 如果 nums[i + 1] < nums[i-1] 那么 我们应该修改nums[i+1]为更大的nums[i] nums[i+1]=nums[i] 这样子就能满足 非递减的数列

// @lc code=start
func checkPossibility(nums []int) bool {
	numsLen := len(nums)
	flag := 0
	for i := 0; i < numsLen-1; i++ {
		if nums[i] > nums[i+1] {
			flag += 1
			if flag == 2 {
				return false
			}
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}

// @lc code=end

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}
