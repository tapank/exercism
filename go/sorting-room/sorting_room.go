package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber returns a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox returns a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber returns the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
	if ok {
		if n, err := strconv.Atoi(fn.Value()); err == nil {
			return n
		}
	}
	return 0
}

// DescribeFancyNumberBox returns a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %d.0", ExtractFancyNumber(fnb))
}

// DescribeAnything returns a string describing whatever it contains.
func DescribeAnything(i interface{}) (s string) {
	switch i.(type) {
	case float64:
		s = DescribeNumber(i.(float64))
	case int:
		s = DescribeNumber(float64(i.(int)))
	case NumberBox:
		s = DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		s = DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		s = "Return to sender"
	}
	return
}
