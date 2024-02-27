package utils

import (
	"time"
)

const (
	DDMMYYYY = "02/01/2006"
)

var (
	weekDayNames = map[time.Weekday]string{
		time.Monday:    "Thứ 2",
		time.Tuesday:   "Thứ 3",
		time.Wednesday: "Thứ 4",
		time.Thursday:  "Thứ 5",
		time.Friday:    "Thứ 6",
		time.Saturday:  "Thứ 7",
		time.Sunday:    "Chủ nhật",
	}
)

// GetYearDiff is a helper function to get the difference in years between two times
func GetYearDiff(end time.Time, start time.Time) int {
	return int(end.Sub(start).Hours() / 24 / 365)
}

func GetNow() time.Time {
	return time.Now()
}

// GetDayDiff is a helper function to get the difference in days between two times
func GetDayDiff(end time.Time, start time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}

// GetWeekdayName returns name of weekday
func GetWeekdayName(t time.Time) string {
	weekday := t.Weekday()
	return weekDayNames[weekday]
}

func ToDateString(t string) string {
	if t != "" {
		// parse string to time
		theTime, _ := time.Parse(time.RFC3339, t)
		// format date only
		return theTime.Format(time.DateOnly)
	}
	return t
}
