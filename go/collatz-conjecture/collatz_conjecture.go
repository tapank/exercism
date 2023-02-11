package collatzconjecture

import "errors"

func CollatzConjecture(n int) (count int, err error) {
	if n < 1 {
		err = errors.New("bad argument")
		return
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		count++
	}
	return
}
