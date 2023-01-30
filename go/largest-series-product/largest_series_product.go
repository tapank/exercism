package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (ls int64, err error) {
	dl := len(digits)
	if dl < span {
		err = errors.New("span cannot be shorter than string")
		return
	}
	if span < 0 {
		err = errors.New("span cannot be negative")
		return
	}
	if dl == 0 {
		ls = 1
		return
	}
	for i := 0; i < dl-span+1; i++ {
		product := int64(1)
		for j := i; j < i+span; j++ {
			if digits[j] >= '0' && digits[j] <= '9' {
				product *= int64(digits[j] - '0')
			} else {
				err = errors.New("invalid digit")
				return
			}
		}
		if product > ls {
			ls = product
		}
	}
	return
}
