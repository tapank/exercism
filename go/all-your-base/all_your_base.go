package allyourbase

import (
	"errors"
)

// ConvertToBase converts input digits of known base to another base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// validations
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// base case
	if inputBase == outputBase {
		return inputDigits, nil
	}

	// convert input to decimal
	decNum, err := otherToDec(inputDigits, inputBase)
	if err != nil {
		return nil, err
	}

	// convert decimal number to target base
	out, ok := decToOther(decNum, outputBase)
	if !ok {
		return nil, errors.New("invalid input")
	}
	return out, nil
}

func decToOther(num, base int) ([]int, bool) {
	if num < 0 {
		return nil, false
	} else if num == 0 {
		return []int{0}, true
	}
	out := []int{}
	for num > 0 {
		out = append(out, num%base)
		num /= base
	}
	// reverse digits
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out, true
}

func otherToDec(num []int, base int) (int, error) {
	decNum := 0
	pow := func(num, exp int) int {
		n := 1
		for i := 0; i < exp; i++ {
			n *= num
		}
		return n
	}
	for i, j := 0, len(num)-1; j >= 0; i, j = i+1, j-1 {
		in := num[i]
		if in < 0 || in >= base {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decNum += pow(base, j) * in
	}
	return decNum, nil
}
