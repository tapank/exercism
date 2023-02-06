package say

var units = map[int]string{
	0:  "",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = map[int]string{
	0: "",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}

	var word string
	// numbers upto thousand
	if word, n = SayInt(int(n%1000)), n/1000; n == 0 {
		return word, true
	}

	// numbers upto million
	if word == "" && n%1000 > 0 {
		word = SayInt(int(n%1000)) + " thousand"
	} else if n%1000 > 0 {
		word = SayInt(int(n%1000)) + " thousand " + word
	}
	if n /= 1000; n == 0 {
		return word, true
	}

	// numbers upto billion
	if word == "" && n%1000 > 0 {
		word = SayInt(int(n%1000)) + " million"
	} else if n%1000 > 0 {
		word = SayInt(int(n%1000)) + " million " + word
	}
	if n /= 1000; n == 0 {
		return word, true
	}

	// numbers upto trillion
	if word == "" {
		word = SayInt(int(n%1000)) + " billion"
	} else {
		word = SayInt(int(n%1000)) + " billion " + word
	}
	return word, true
}

// SayInt works for numbers upto thousand (not inclusive)
func SayInt(n int) string {
	var word string
	if n%100 > 0 && n%100 < 20 {
		word = units[n%100]
		n /= 100
	} else {
		word = units[n%10]
		n /= 10
		if n == 0 {
			return word
		}
		ten := tens[n%10]
		if word != "" {
			word = ten + "-" + word
		} else {
			word = ten
		}
	}
	if n /= 10; n == 0 {
		return word
	}
	hundred := units[n%10]
	if hundred == "" {
		return word
	}
	if word == "" {
		word = hundred + " hundred"
	} else {
		word = hundred + " hundred " + word
	}
	return word
}
