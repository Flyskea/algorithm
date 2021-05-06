/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */
package main

import "fmt"

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	var res string
	sLen := len(s)
	dictionaryLen := len(dictionary)
	for k := 0; k < dictionaryLen; k++ {
		i, j := 0, 0
		v := dictionary[k]
		vLen := len(v)
		for i < sLen && j < vLen {
			if s[i] == v[j] {
				j++
			}
			i++
		}
		if j == vLen && (len(res) < vLen || (len(res) == vLen) && res > v) {
			res = v
		}
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}
