package main

import (
	"fmt"
)

func main() {

	n := 0
	fmt.Println(n, digitCount(n))
	n = 1
	fmt.Println(n, digitCount(n))
	n = 99
	fmt.Println(n, digitCount(n))
	n = 100
	fmt.Println(n, digitCount(n))
	n = 999
	fmt.Println(n, digitCount(n))
	n = 1000
	fmt.Println(n, digitCount(n))
	n = 1001
	fmt.Println(n, digitCount(n))
	n = -333
	fmt.Println(n, digitCount(n))
}

func digitCount(n int) (cnt int) {
	for n > 0 {
		cnt++
		n /= 10
	}
	return
}
