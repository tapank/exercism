package ledger

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var headerFmt = "%-10s | %-25s | %s\n"
var headerFields = map[string][]string{
	"nl-NL": {"Datum", "Omschrijving", "Verandering"},
	"en-US": {"Date", "Description", "Change"},
}

// currenty symbols
var currencySymbols = map[string]string{
	"EUR": "€",
	"USD": "$",
}

// all the separators needed to format an amount by locale
type NumFormat struct {
	prefix        string
	tsep, decimal rune
}

// the amount number separators for locales
var separators = map[string]NumFormat{
	"nl-NL": {" ", '.', ','},
	"en-US": {"", ',', '.'},
}

// date formats for locales.
var dateFormat = map[string]string{
	"nl-NL": "02-01-2006",
	"en-US": "01/02/2006",
}

// FormatLedger formats given entries based on given locale and currency.
// To add support for new currencies and locales, update the date formats,
// separators, currency symbols.
func FormatLedger(currency string, locale string, entries []Entry) (out string, err error) {
	// get header. this also validates locale.
	if out, err = header(locale); err != nil {
		return
	}

	// validate currency
	if _, ok := currencySymbols[currency]; !ok {
		if currency == "" {
			err = errors.New("empty currency")
		} else {
			err = errors.New("invalid currency")
		}
		return
	}

	// create a copy in order to not modify the input array
	es := append([]Entry{}, entries...)

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

	// format each entry and add to output string
	for _, entry := range es {
		// parse and format date
		var datestr string
		var t time.Time
		if t, err = time.Parse("2006-01-02", entry.Date); err != nil {
			return
		} else {
			datestr = t.Format(dateFormat[locale])
		}

		description := entry.Description
		if len(description) > 25 {
			description = description[:22] + "..."
		}
		var amount string
		if amount, err = formatCurrency(entry.Change, currency, locale); err != nil {
			return
		}
		out += fmt.Sprintf("%10s | %-25s | %13s\n", datestr, description, amount)
	}
	return
}

// header creates a header for a given locale
func header(locale string) (out string, err error) {
	if fields, ok := headerFields[locale]; ok {
		out = fmt.Sprintf(headerFmt, fields[0], fields[1], fields[2])
		return
	}
	err = errors.New("unsupported locale")
	return
}

// formatCurrency stringifies amount based on locale and currency.
// handling of negative amounts needs improvement (I am not proud).
func formatCurrency(cents int, currency, locale string) (out string, err error) {
	// establish the number formatting separators
	var tsep, dec rune
	var prefix string
	if numSep, ok := separators[locale]; ok {
		tsep, dec = numSep.tsep, numSep.decimal
		prefix = numSep.prefix
	} else {
		err = errors.New("invalid locale")
		return
	}

	// remember the sign and make cents positive
	var isNeg bool
	if cents < 0 {
		isNeg = true
		cents = -cents
	}

	// now convert the amount to string
	units, change := cents/100, cents%100
	out = fmt.Sprintf("%c%02d", dec, change)
	for units/1000 > 0 {
		out = fmt.Sprintf("%c%03d%s", tsep, units%1000, out)
		units /= 1000
	}
	out = fmt.Sprintf("%s%s%d%s", currencySymbols[currency], prefix, units, out)
	if isNeg {
		if tsep == ',' {
			out = "(" + out + ")"
		} else {
			out += "-"
		}
	} else {
		out += " "
	}
	return
}
