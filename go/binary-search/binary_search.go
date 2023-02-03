package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 || key < list[0] || key > list[len(list)-1] {
		return -1
	}
	l, r := 0, len(list)
	m := (r - l) / 2
	for m >= l && m <= r {
		// found match
		if list[m] == key {
			return m
		}

		// divide range
		if list[m] < key {
			l = m
		} else if list[m] > key {
			r = m
		}

		// find new m
		nextm := l + (r-l)/2
		if m == nextm {
			break
		}
		m = nextm
	}
	return -1
}
