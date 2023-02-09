package bob

import "strings"

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	isQ, isY := isQuestion(remark), isYell(remark)
	if isQ && isY {
		return "Calm down, I know what I'm doing!"
	}
	if isQ {
		return "Sure."
	}
	if isY {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[len(s)-1] == '?'
}

func isYell(s string) bool {
	var yell bool
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return false
		}
		if r >= 'A' && r <= 'Z' {
			yell = true
		}
	}
	return yell
}
