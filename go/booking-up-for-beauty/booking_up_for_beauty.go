package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// "July 25, 2019 13:45:00"
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// "Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// input "7/25/2019 13:45:00"
	// output "You have an appointment on Thursday, July 25, 2019, at 13:45."
	inlay := "1/2/2006 15:04:05"
	outlay := "You have an appointment on Monday, January 2, 2006, at 15:04."
	t, _ := time.Parse(inlay, date)
	return t.Format(outlay)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// 2020-09-15 00:00:00 +0000 UTC
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.Now().UTC().Location())
}
