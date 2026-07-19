package timex

import (
	"fmt"
	"time"
)

const (
	LayoutDate     = "2006-01-02"
	LayoutTime     = "15:04:05"
	LayoutDateTime = "2006-01-02 15:04:05"
	LayoutISO8601  = "2006-01-02T15:04:05Z07:00"
	LayoutRFC3339  = time.RFC3339
)

// FormatDate formats time to date string (2006-01-02).
func FormatDate(t time.Time) string {
	return t.Format(LayoutDate)
}

// FormatTime formats time to time string (15:04:05).
func FormatTime(t time.Time) string {
	return t.Format(LayoutTime)
}

// FormatDateTime formats time to datetime string (2006-01-02 15:04:05).
func FormatDateTime(t time.Time) string {
	return t.Format(LayoutDateTime)
}

// FormatISO8601 formats time to ISO8601 string.
func FormatISO8601(t time.Time) string {
	return t.Format(LayoutISO8601)
}

// ParseDate parses date string.
func ParseDate(s string) (time.Time, error) {
	return time.Parse(LayoutDate, s)
}

// ParseDateTime parses datetime string.
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse(LayoutDateTime, s)
}

// ParseISO8601 parses ISO8601 string.
func ParseISO8601(s string) (time.Time, error) {
	return time.Parse(LayoutISO8601, s)
}

// MustParseDateTime parses datetime, panics on error.
func MustParseDateTime(s string) time.Time {
	t, err := ParseDateTime(s)
	if err != nil {
		panic(err)
	}
	return t
}

// MustParseDate parses date, panics on error.
func MustParseDate(s string) time.Time {
	t, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return t
}

// Parse 自动识别常见格式
func Parse(s string) (time.Time, error) {
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006/01/02 15:04:05",
		"2006/01/02",
		time.RFC3339,
	}

	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse time: %s", s)
}
