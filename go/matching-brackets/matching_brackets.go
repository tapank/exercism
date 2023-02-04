package brackets

var pair = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

func Bracket(input string) bool {
	stack := make([]rune, 0, 10)
	for _, r := range input {
		switch r {
		case '[', '{', '(':
			stack = append(stack, r)
		case ']', '}', ')':
			if l := len(stack); l == 0 || stack[l-1] != pair[r] {
				return false
			} else {
				stack = stack[:l-1]
			}
		}
	}
	return len(stack) == 0
}
