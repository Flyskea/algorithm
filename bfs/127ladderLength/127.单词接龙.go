/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */
package main

import "fmt"

// @lc code=start

func getCandidates(word string) []string {
	var res []string
	wordLen := len(word)
	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j < wordLen; j++ {
			candidate := i
			if word[j] != byte(candidate) {
				res = append(res, word[:j]+string(candidate)+word[j+1:])
			}
		}
	}
	return res
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, que, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0
	for len(que) > 0 {
		depth++
		qLen := len(que)
		for i := 0; i < qLen; i++ {
			word := que[0]
			que = que[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}
			}
		}
	}
	return 0
}

// @lc code=end

/*
双向遍历
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if wordList == nil || len(wordList) == 0 {
		return 0
	}
	if endWord == "" || beginWord == "" {
		return 0
	}

	myMap := make(map[string]bool, len(wordList))
	for _, str := range wordList {
		myMap[str] = true
	}
	if _, ok := myMap[endWord]; !ok {
		return 0
	}
	delete(myMap, beginWord)

	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	beginVisited, endVisited := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	step := 1

	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}

		newVisited := map[string]bool{}
		for word := range beginVisited {
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					if rune(word[i]) == ch {
						continue
					}
					newWord := word[:i] + string(ch) + word[i+1:]
					if _, ok := myMap[newWord]; ok {
						if _, ok := endVisited[newWord]; ok {
							step++
							return step
						}

						if _, yes := visited[newWord]; !yes {
							visited[newWord] = true
							newVisited[newWord] = true
						}
					}
				}
			}
		}
		beginVisited = newVisited
		step++
	}
*/

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
