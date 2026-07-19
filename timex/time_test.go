package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	tm := time.Date(2026, 7, 19, 10, 30, 0, 0, time.Local)
	assert.Equal(t, "2026-07-19", FormatDate(tm))
}

func TestFormatDateTime(t *testing.T) {
	tm := time.Date(2026, 7, 19, 10, 30, 0, 0, time.Local)
	assert.Equal(t, "2026-07-19 10:30:00", FormatDateTime(tm))
}

func TestParseDate(t *testing.T) {
	tm, err := ParseDate("2026-07-19")
	assert.NoError(t, err)
	assert.Equal(t, 2026, tm.Year())
	assert.Equal(t, time.July, tm.Month())
	assert.Equal(t, 19, tm.Day())
}

func TestParseDateTime(t *testing.T) {
	tm, err := ParseDateTime("2026-07-19 10:30:00")
	assert.NoError(t, err)
	assert.Equal(t, 10, tm.Hour())
	assert.Equal(t, 30, tm.Minute())
}

func TestMustParseDateTime(t *testing.T) {
	tm := MustParseDateTime("2026-07-19 10:30:00")
	assert.Equal(t, 2026, tm.Year())
	assert.Panics(t, func() { MustParseDateTime("invalid") })
}

func TestStartOfDay(t *testing.T) {
	tm := time.Date(2026, 7, 19, 10, 30, 45, 100, time.Local)
	start := StartOfDay(tm)
	assert.Equal(t, 0, start.Hour())
	assert.Equal(t, 0, start.Minute())
	assert.Equal(t, 0, start.Second())
}

func TestEndOfDay(t *testing.T) {
	tm := time.Date(2026, 7, 19, 10, 30, 0, 0, time.Local)
	end := EndOfDay(tm)
	assert.Equal(t, 23, end.Hour())
	assert.Equal(t, 59, end.Minute())
	assert.Equal(t, 59, end.Second())
}

func TestAddDays(t *testing.T) {
	tm := time.Date(2026, 7, 19, 0, 0, 0, 0, time.Local)
	next := AddDays(tm, 1)
	assert.Equal(t, 20, next.Day())
}

func TestIsWeekend(t *testing.T) {
	sat := time.Date(2026, 7, 18, 0, 0, 0, 0, time.Local)
	sun := time.Date(2026, 7, 19, 0, 0, 0, 0, time.Local)
	mon := time.Date(2026, 7, 20, 0, 0, 0, 0, time.Local)

	assert.True(t, IsWeekend(sat))
	assert.True(t, IsWeekend(sun))
	assert.False(t, IsWeekend(mon))
}

func TestBetween(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.Local)
	mid := time.Date(2026, 6, 15, 0, 0, 0, 0, time.Local)

	assert.True(t, Between(mid, start, end))
	assert.True(t, Between(start, start, end))
	assert.False(t, Between(end.Add(time.Hour), start, end))
}

func TestTimestamp(t *testing.T) {
	tm := time.Date(2026, 7, 19, 0, 0, 0, 0, time.UTC)
	ts := Timestamp(tm)
	assert.Greater(t, ts, int64(0))
}

func TestYesterday(t *testing.T) {
	y := Yesterday()
	assert.Equal(t, StartOfDay(time.Now().AddDate(0, 0, -1)), y)
}

func TestTomorrow(t *testing.T) {
	tm := Tomorrow()
	assert.Equal(t, StartOfDay(time.Now().AddDate(0, 0, 1)), tm)
}

func TestIsMonday(t *testing.T) {
	mon := time.Date(2026, 7, 20, 0, 0, 0, 0, time.Local)
	sun := time.Date(2026, 7, 19, 0, 0, 0, 0, time.Local)

	assert.True(t, IsMonday(mon))
	assert.False(t, IsMonday(sun))
}

func TestStartOfWeek(t *testing.T) {
	// 2026-07-19 周日
	sun := time.Date(2026, 7, 19, 10, 0, 0, 0, time.Local)
	start := StartOfWeek(sun)
	// 应该是 2026-07-13 周一
	assert.Equal(t, 13, start.Day())
	assert.Equal(t, time.Monday, start.Weekday())

	// 2026-07-20 周一
	mon := time.Date(2026, 7, 20, 10, 0, 0, 0, time.Local)
	start = StartOfWeek(mon)
	assert.Equal(t, 20, start.Day())
}

func TestEndOfWeek(t *testing.T) {
	// 2026-07-20 周一
	mon := time.Date(2026, 7, 20, 10, 0, 0, 0, time.Local)
	end := EndOfWeek(mon)
	// 应该是 2026-07-26 周日
	assert.Equal(t, 26, end.Day())
	assert.Equal(t, 23, end.Hour())
}

func TestNextMonday(t *testing.T) {
	// 2026-07-20 周一
	mon := time.Date(2026, 7, 20, 10, 0, 0, 0, time.Local)
	next := NextMonday(mon)
	// 应该是 2026-07-27
	assert.Equal(t, 27, next.Day())

	// 2026-07-19 周日
	sun := time.Date(2026, 7, 19, 10, 0, 0, 0, time.Local)
	next = NextMonday(sun)
	assert.Equal(t, 20, next.Day())
}

func TestPrevMonday(t *testing.T) {
	// 2026-07-20 周一
	mon := time.Date(2026, 7, 20, 10, 0, 0, 0, time.Local)
	prev := PrevMonday(mon)
	// 应该是 2026-07-13
	assert.Equal(t, 13, prev.Day())
}

func TestWeekdayName(t *testing.T) {
	mon := time.Date(2026, 7, 20, 0, 0, 0, 0, time.Local)
	assert.Equal(t, "周一", WeekdayName(mon))

	sun := time.Date(2026, 7, 19, 0, 0, 0, 0, time.Local)
	assert.Equal(t, "周日", WeekdayName(sun))
}

func TestDaysInMonth(t *testing.T) {
	jan := time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, 31, DaysInMonth(jan))

	feb := time.Date(2026, 2, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, 28, DaysInMonth(feb))

	leapFeb := time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, 29, DaysInMonth(leapFeb))
}
