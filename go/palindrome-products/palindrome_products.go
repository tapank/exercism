package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Val            int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	if fmin < 0 {
		return Product{}, Product{}, errors.New("illegal limits")
	}
	palindromes := map[int]*Product{}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			factors := [2]int{i, j}
			if p, ok := palindromes[product]; ok {
				p.Factorizations = append(p.Factorizations, factors)
			} else if isPalindrome(product) {
				palindromes[product] = &Product{product, [][2]int{factors}}
			}
		}
	}
	if len(palindromes) == 0 {
		return Product{}, Product{}, errors.New("no palindromes")
	}
	min, max := -1, -1
	for k := range palindromes {
		if min == -1 || k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}
	return *palindromes[min], *palindromes[max], nil
}

func isPalindrome(n int) bool {
	numstr := strconv.Itoa(n)
	if len(numstr) == 1 {
		return true
	}
	for i, j := 0, len(numstr)-1; i < j; i, j = i+1, j-1 {
		if numstr[i] != numstr[j] {
			return false
		}
	}
	return true
}
