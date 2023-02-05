package wordy

import (
	"strconv"
	"strings"
)

// supported operations
var operations = map[string]func(op1, op2 int) int{
	"plus":          func(op1, op2 int) int { return op1 + op2 },
	"minus":         func(op1, op2 int) int { return op1 - op2 },
	"multiplied by": func(op1, op2 int) int { return op1 * op2 },
	"divided by":    func(op1, op2 int) int { return op1 / op2 },
}

func Answer(question string) (int, bool) {
	// non math question.
	if !strings.HasPrefix(question, "What is ") {
		return 0, false
	}

	// trim cruft at both ends.
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimRight(question, "?.!")
	question = strings.TrimSpace(question)

	// missing expression
	if len(question) == 0 {
		return 0, false
	}

	var op1, op2 int
	var operator string
	for i, token := range strings.Split(question, " ") {
		var err error
		// results accumulate in op1
		if i == 0 {
			if op1, err = strconv.Atoi(token); err != nil {
				return 0, false
			}
			continue
		}

		switch token {
		case "plus", "minus", "multiplied", "divided":
			if operator != "" {
				// found repeated operation
				return 0, false
			}
			operator = token
		case "by":
			operator += " by"
		default:
			if op2, err = strconv.Atoi(token); err != nil {
				return 0, false
			}
			if f, ok := operations[operator]; ok {
				op1 = f(op1, op2)
				operator = ""
			} else {
				// operation missing or unsupported operation
				return 0, false
			}
		}
	}
	// missing operand
	if operator != "" {
		return 0, false
	}
	return op1, true
}
