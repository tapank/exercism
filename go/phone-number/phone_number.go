package phonenumber

import (
	"errors"
	"fmt"
)

// Parse returns area code and number; or error if any validation fails
func Parse(pnum string) (acode, num string, err error) {
	if len(pnum) < 10 {
		err = errors.New("not enough digits")
		return
	}

	// trim leading + and country code if present
	if pnum[0] == '+' {
		pnum = pnum[1:]
	}
	if pnum[0] == '1' {
		pnum = pnum[1:]
	}

	// extract area code and number
	dcount := 0
	for _, r := range pnum {
		if r >= '0' && r <= '9' {
			// 0 and 1 not allowed in the 0 and 3 index
			if (dcount == 0 || dcount == 3) && r < '2' {
				err = errors.New("illegal starting of area code or number")
				return
			}

			if dcount < 3 {
				acode += string(r)
			} else if dcount < 10 {
				num += string(r)
			} else {
				err = errors.New("too many digits")
				return
			}
			dcount++
		}
		// ignore other characters
	}

	// area code and number should be 10 digits
	if len(acode) != 3 || len(num) != 7 {
		err = errors.New("not enough digits")
		return
	}
	return
}

// Number parses given phone number and returns stripped out
// 10 digit number excluding leading +1 and other formatting.
func Number(phoneNumber string) (string, error) {
	acode, num, err := Parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return acode + num, nil
}

// AreaCode parses given phone number and returns stripped out
// 3 digit area code.
func AreaCode(phoneNumber string) (string, error) {
	acode, _, err := Parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return acode, nil
}

// Format returns the phone number in "(NXX) NXX-XXXX" format.
func Format(phoneNumber string) (string, error) {
	acode, num, err := Parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", acode, num[0:3], num[3:]), nil
}
