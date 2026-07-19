package timex

import "time"

// Now returns current time.
func Now() time.Time {
	return time.Now()
}

// Today returns start of today.
func Today() time.Time {
	return StartOfDay(time.Now())
}

// StartOfDay returns start of day (00:00:00).
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns end of day (23:59:59.999999999).
func EndOfDay(t time.Time) time.Time {
	return StartOfDay(t).Add(24*time.Hour - time.Nanosecond)
}

// StartOfMonth returns start of month.
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns end of month.
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// StartOfYear returns start of year.
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns end of year.
func EndOfYear(t time.Time) time.Time {
	return StartOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// AddDays adds days to time.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths adds months to time.
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds years to time.
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// SubDays returns difference in days.
func SubDays(t1, t2 time.Time) int {
	return int(t1.Sub(t2).Hours() / 24)
}

// IsWeekend checks if time is weekend.
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	return wd == time.Saturday || wd == time.Sunday
}

// IsToday checks if time is today.
func IsToday(t time.Time) bool {
	return FormatDate(t) == FormatDate(time.Now())
}

// Age calculates age from birth date.
func Age(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

// Between checks if t is between start and end.
func Between(t, start, end time.Time) bool {
	return (t.Equal(start) || t.After(start)) && (t.Equal(end) || t.Before(end))
}

// Timestamp returns Unix timestamp in seconds.
func Timestamp(t time.Time) int64 {
	return t.Unix()
}

// TimestampMilli returns Unix timestamp in milliseconds.
func TimestampMilli(t time.Time) int64 {
	return t.UnixMilli()
}
