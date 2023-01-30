package meetup

import "time"

type WeekSchedule int

const (
	First = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day returns the day of the month based on WeekSchedule in a given month and year
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// time on 1st of the given month and year
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// number of days to the next wDay, i.e. the first wDay of the month
	daysToWDay := int(wDay - t.Weekday())
	if daysToWDay < 0 {
		daysToWDay += 7
	}

	// move time forward to the first wDay of the given month
	// remaining calculations use this value as base
	t = t.AddDate(0, 0, daysToWDay)
	switch wSched {
	case First:
		// no-op, t is already on first wDay
	case Second:
		// add one week
		t = t.AddDate(0, 0, 7)
	case Third:
		// add two weeks
		t = t.AddDate(0, 0, 14)
	case Fourth:
		// add three weeks
		t = t.AddDate(0, 0, 21)
	case Teenth:
		// start with second wDay
		t = t.AddDate(0, 0, 7)
		// check if second wDay is a teenth, and add a week if not
		if t.Day() < 13 {
			t = t.AddDate(0, 0, 7)
		}
	case Last:
		// start with 4 weeks after first wDay
		t = t.AddDate(0, 0, 28)
		// if month has rolled over, then move back a week
		if t.Month() != month {
			t = t.AddDate(0, 0, -7)
		}
	}
	return t.Day()
}
