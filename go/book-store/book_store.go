package bookstore

func Cost(books []int) int {
	cart := map[int]int{}
	var amount int
	for _, book := range books {
		cart[book]++
	}
	reduceOneOfEach := func(cart map[int]int) {
		for k := range cart {
			cart[k]--
			if cart[k] == 0 {
				delete(cart, k)
			}
		}
	}

	for len(cart) == 5 {
		reduceOneOfEach(cart)
		amount += 5 * 8 * 75
	}
	for len(cart) == 4 {
		reduceOneOfEach(cart)
		amount += 4 * 8 * 80
	}
	for len(cart) == 3 {
		reduceOneOfEach(cart)
		amount += 3 * 8 * 90
	}
	for len(cart) == 2 {
		reduceOneOfEach(cart)
		amount += 2 * 8 * 95
	}
	for len(cart) == 1 {
		reduceOneOfEach(cart)
		amount += 1 * 8 * 100
	}
	return amount
}
