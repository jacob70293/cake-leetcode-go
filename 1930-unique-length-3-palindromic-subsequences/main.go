package _1930_unique_length_3_palindromic_subsequences

import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {

	// 先把字母所有的 index 找出來
	letterIndexesMap := make(map[rune][]int)
	for i, r := range s {
		letterIndexesMap[r] = append(letterIndexesMap[r], i)
	}

	// 回傳組合 (要排除重複)
	resultAssoc := map[string]bool{}

	// 遍歷所有 index slice，頭尾取中間字母做組合
	for r, indexes := range letterIndexesMap {
		// 如果 index count < 2，就不用掃中間值了
		if len(indexes) < 2 {
			continue
		}

		// 抓出最初 & 最後字母位置
		start, end := indexes[0], indexes[len(indexes)-1]

		// 中間沒有其他字母
		if end-start < 2 {
			continue
		}

		for _, r2 := range s[start+1 : end] {
			result := fmt.Sprintf("%c%c%c", r, r2, r)
			resultAssoc[result] = true
		}
	}

	return len(resultAssoc)
}
