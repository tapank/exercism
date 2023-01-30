package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (keep Ints) {
	for _, n := range i {
		if filter(n) {
			keep = append(keep, n)
		}
	}
	return
}

func (i Ints) Discard(filter func(int) bool) (keep Ints) {
	for _, n := range i {
		if !filter(n) {
			keep = append(keep, n)
		}
	}
	return
}

func (l Lists) Keep(filter func([]int) bool) (keep Lists) {
	for _, s := range l {
		if filter(s) {
			keep = append(keep, s)
		}
	}
	return
}

func (s Strings) Keep(filter func(string) bool) (keep Strings) {
	for _, v := range s {
		if filter(v) {
			keep = append(keep, v)
		}
	}
	return
}
