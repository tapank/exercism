package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return s
	}
	new := IntList{}
	for _, v := range s {
		if fn(v) {
			new = append(new, v)
		}
	}
	return new
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	if len(s) == 0 {
		return s
	}
	new := make(IntList, len(s))
	for i, v := range s {
		new[i] = fn(v)
	}
	return new
}

func (s IntList) Reverse() IntList {
	if len(s) == 0 {
		return s
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	if s == nil {
		s = make(IntList, 0, len(lst))
	}
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}
	return s
}
