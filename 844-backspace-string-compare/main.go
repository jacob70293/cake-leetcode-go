package _844_backspace_string_compare

func backspaceCompare(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	var stack []rune
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
