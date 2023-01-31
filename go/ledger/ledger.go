package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (out string, err error) {
	var s string

	if currency == "" {
		err = errors.New("empty currency")
		return
	}
	if currency != "EUR" && currency != "USD" {
		err = errors.New("invalid currency")
		return
	}

	// header
	headerFmt := "%-10s | %-25s | %s\n"
	switch locale {
	case "nl-NL":
		s = fmt.Sprintf(headerFmt, "Datum", "Omschrijving", "Verandering")
	case "en-US":
		s = fmt.Sprintf(headerFmt, "Date", "Description", "Change")
	default:
		err = errors.New("unsupported locale")
		return
	}

	// create a copy in order to not modify the input array
	var es []Entry
	es = append(es, entries...)

	// sort entries by date, description, and change
	sort.Slice(es, func(i, j int) bool {
		ies, jes := es[i], es[j]
		if ies.Date != jes.Date {
			return ies.Date < jes.Date
		}
		if ies.Description != jes.Description {
			return ies.Description < jes.Description
		}
		return ies.Change < jes.Change
	})

	ss := make([]string, len(es))
	for i, entry := range es {
		if len(entry.Date) != 10 {
			return "", errors.New("")
		}
		d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
		if d2 != '-' {
			return "", errors.New("")
		}
		if d4 != '-' {
			return "", errors.New("")
		}
		de := entry.Description
		if len(de) > 25 {
			de = de[:22] + "..."
		} else {
			de = de + strings.Repeat(" ", 25-len(de))
		}
		var d string
		if locale == "nl-NL" {
			d = d5 + "-" + d3 + "-" + d1
		} else if locale == "en-US" {
			d = d3 + "/" + d5 + "/" + d1
		}
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}
		var a string
		if locale == "nl-NL" {
			if currency == "EUR" {
				a += "€"
			} else if currency == "USD" {
				a += "$"
			}
			a += " "
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + "."
			}
			a = a[:len(a)-1]
			a += ","
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += "-"
			} else {
				a += " "
			}
		} else if locale == "en-US" {
			if negative {
				a += "("
			}
			if currency == "EUR" {
				a += "€"
			} else if currency == "USD" {
				a += "$"
			}
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + ","
			}
			a = a[:len(a)-1]
			a += "."
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += ")"
			} else {
				a += " "
			}
		}
		var al int
		for range a {
			al++
		}
		ss[i] = d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
			strings.Repeat(" ", 13-al) + a + "\n"
	}
	for i := 0; i < len(es); i++ {
		s += ss[i]
	}
	fmt.Println(s)
	return s, nil
}
