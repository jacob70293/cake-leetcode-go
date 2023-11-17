package _20_valid_parentheses

import (
	"slices"
)

const (
	LeftParenthesis  rune = '('
	RightParenthesis rune = ')'
	LeftBracket      rune = '['
	RightBracket     rune = ']'
	LeftBrace        rune = '{'
	RightBrace       rune = '}'
)

var pairParentheses map[rune]rune

func init() {
	pairParentheses = map[rune]rune{
		LeftParenthesis: RightParenthesis,
		LeftBracket:     RightBracket,
		LeftBrace:       RightBrace,
	}
}

func isValid(s string) bool {
	var runeStack []rune

	for _, r := range s {
		// 符合左括號，放進暫存列表
		if slices.Contains([]rune{LeftParenthesis, LeftBracket, LeftBrace}, r) {
			runeStack = append(runeStack, r)
			continue
		}

		if len(runeStack) == 0 {
			return false
		}

		leftRune := runeStack[len(runeStack)-1]

		// pop stack 比對
		if pairParentheses[leftRune] != r {
			return false
		}

		runeStack = runeStack[0 : len(runeStack)-1]
	}

	return len(runeStack) == 0
}
