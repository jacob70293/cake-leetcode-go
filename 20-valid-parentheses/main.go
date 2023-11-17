package _20_valid_parentheses

var leftParentheses = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
}

var pairParentheses = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isValid(s string) bool {
	var tempStack []string

	for _, r := range s {
		// 符合左括號，放進暫存列表
		if leftParentheses[string(r)] {
			tempStack = append(tempStack, string(r))
		} else { // 右括號
			if len(tempStack) == 0 {
				return false
			}

			// pop stack 比對
			popEle := tempStack[len(tempStack)-1]
			if pairParentheses[string(r)] != popEle {
				return false
			}

			tempStack = tempStack[0 : len(tempStack)-1]
		}
	}

	if len(tempStack) > 0 {
		return false
	}

	return true
}
