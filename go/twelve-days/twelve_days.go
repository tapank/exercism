package twelve

import "fmt"

var dayNames = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	prefix := func(d int) string {
		return fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", dayNames[i])
	}
	suffix := ""
	switch i {
	case 12:
		suffix += "twelve Drummers Drumming, "
		fallthrough
	case 11:
		suffix += "eleven Pipers Piping, "
		fallthrough
	case 10:
		suffix += "ten Lords-a-Leaping, "
		fallthrough
	case 9:
		suffix += "nine Ladies Dancing, "
		fallthrough
	case 8:
		suffix += "eight Maids-a-Milking, "
		fallthrough
	case 7:
		suffix += "seven Swans-a-Swimming, "
		fallthrough
	case 6:
		suffix += "six Geese-a-Laying, "
		fallthrough
	case 5:
		suffix += "five Gold Rings, "
		fallthrough
	case 4:
		suffix += "four Calling Birds, "
		fallthrough
	case 3:
		suffix += "three French Hens, "
		fallthrough
	case 2:
		suffix += "two Turtle Doves, and "
		fallthrough
	case 1:
		suffix += "a Partridge in a Pear Tree."
	}
	return prefix(i) + suffix
}

func Song() string {
	song := Verse(1)
	for i := 2; i <= 12; i++ {
		song += "\n" + Verse(i)
	}
	return song
}
