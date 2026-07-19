package timex

import "time"

// Yesterday returns yesterday's date at start of day.
func Yesterday() time.Time {
	return StartOfDay(time.Now().AddDate(0, 0, -1))
}

// Tomorrow returns tomorrow's date at start of day.
func Tomorrow() time.Time {
	return StartOfDay(time.Now().AddDate(0, 0, 1))
}

// DayOfWeek represents Monday to Sunday.
type DayOfWeek int

const (
	Monday    DayOfWeek = 1
	Tuesday   DayOfWeek = 2
	Wednesday DayOfWeek = 3
	Thursday  DayOfWeek = 4
	Friday    DayOfWeek = 5
	Saturday  DayOfWeek = 6
	Sunday    DayOfWeek = 0
)

// Weekday returns the DayOfWeek for given time.
func Weekday(t time.Time) DayOfWeek {
	return DayOfWeek(t.Weekday())
}

// IsMonday checks if time is Monday.
func IsMonday(t time.Time) bool {
	return t.Weekday() == time.Monday
}

// IsTuesday checks if time is Tuesday.
func IsTuesday(t time.Time) bool {
	return t.Weekday() == time.Tuesday
}

// IsWednesday checks if time is Wednesday.
func IsWednesday(t time.Time) bool {
	return t.Weekday() == time.Wednesday
}

// IsThursday checks if time is Thursday.
func IsThursday(t time.Time) bool {
	return t.Weekday() == time.Thursday
}

// IsFriday checks if time is Friday.
func IsFriday(t time.Time) bool {
	return t.Weekday() == time.Friday
}

// IsSaturday checks if time is Saturday.
func IsSaturday(t time.Time) bool {
	return t.Weekday() == time.Saturday
}

// IsSunday checks if time is Sunday.
func IsSunday(t time.Time) bool {
	return t.Weekday() == time.Sunday
}

// StartOfWeek returns start of week (Monday 00:00:00).
func StartOfWeek(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	return StartOfDay(t.AddDate(0, 0, -wd+1))
}

// EndOfWeek returns end of week (Sunday 23:59:59).
func EndOfWeek(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	return EndOfDay(t.AddDate(0, 0, 7-wd))
}

// NextMonday returns next Monday.
func NextMonday(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	return StartOfDay(t.AddDate(0, 0, 8-wd))
}

// PrevMonday returns previous Monday.
func PrevMonday(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	return StartOfDay(t.AddDate(0, 0, -wd-6))
}

// WeekdayName returns Chinese weekday name.
func WeekdayName(t time.Time) string {
	names := []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	return names[t.Weekday()]
}

// WeekdayNameEN returns English weekday name.
func WeekdayNameEN(t time.Time) string {
	return t.Weekday().String()
}

// DaysInMonth returns number of days in month.
func DaysInMonth(t time.Time) int {
	// 当月1号
	first := StartOfMonth(t)
	// 下月1号减1天 = 当月最后一天
	last := first.AddDate(0, 1, -1)
	return last.Day()
}
