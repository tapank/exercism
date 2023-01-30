package clock

import "fmt"

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	if m < 0 {
		h -= (-m)/60 + 1
		m = m%60 + 60
	}
	h = (h + m/60) % 24
	m %= 60
	if h < 0 {
		h += (((-h) / 24) + 1) * 24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
