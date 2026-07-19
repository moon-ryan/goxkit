# goxkit

A collection of commonly used Go utility packages, providing lightweight extensions for time handling, random number generation, data type conversion, and more.

## Installation

```bash
go get github.com/moon-ryan/goxkit
```

## Modules

| Package | Description |
| --- | --- |
| [timex](#timex) | Time formatting and calculation utilities |
| [randx](#randx) | Random number and string generation |
| [convert](#convert) | Data type conversion utilities |

## timex

Time handling utilities that wrap common date/time formatting and calculation methods.

### Constants

| Constant | Format |
| --- | --- |
| `LayoutDate` | `2006-01-02` |
| `LayoutTime` | `15:04:05` |
| `LayoutDateTime` | `2006-01-02 15:04:05` |
| `LayoutISO8601` | `2006-01-02T15:04:05Z07:00` |
| `LayoutRFC3339` | `time.RFC3339` |

### Formatting

```go
t := time.Now()
timex.FormatDate(t)     // 2026-07-19
timex.FormatTime(t)     // 10:30:00
timex.FormatDateTime(t) // 2026-07-19 10:30:00
timex.FormatISO8601(t)  // 2026-07-19T10:30:00+08:00
timex.WeekdayName(t)    // Saturday (Chinese locale returns "周六")
timex.WeekdayNameEN(t)  // Saturday
```

### Parsing

```go
t, err := timex.ParseDate("2026-07-19")
t, err = timex.ParseDateTime("2026-07-19 10:30:00")
t, err = timex.ParseISO8601("2026-07-19T10:30:00+08:00")
t = timex.MustParseDateTime("2026-07-19 10:30:00")

// Automatically detect common formats
t, err = timex.Parse("2026-07-19 10:30")
```

### Date Calculation

```go
now := time.Now()

timex.Today()
timex.Yesterday()
timex.Tomorrow()

timex.StartOfDay(now)
timex.EndOfDay(now)
timex.StartOfWeek(now)   // Monday 00:00:00 of the current week
timex.EndOfWeek(now)     // Sunday 23:59:59 of the current week
timex.StartOfMonth(now)
timex.EndOfMonth(now)
timex.StartOfYear(now)
timex.EndOfYear(now)

timex.AddDays(now, 7)
timex.AddMonths(now, 1)
timex.AddYears(now, 1)

timex.NextMonday(now)
timex.PrevMonday(now)

timex.DaysInMonth(now)
timex.SubDays(t1, t2)
timex.Age(birthDate)
```

### Checks

```go
timex.IsWeekend(t)
timex.IsToday(t)
timex.IsMonday(t)
timex.IsTuesday(t)
timex.IsWednesday(t)
timex.IsThursday(t)
timex.IsFriday(t)
timex.IsSaturday(t)
timex.IsSunday(t)
timex.Between(t, start, end)
```

### Timestamps

```go
timex.Timestamp(t)
timex.TimestampMilli(t)
```

## randx

Random number generation utilities, based on the standard library `math/rand`.

```go
randx.Int(100)              // [0, 100)
randx.IntRange(10, 20)      // [10, 20)
randx.Float64()             // [0.0, 1.0)
randx.Bool()                // true / false

randx.String(10)            // 10-character random alphanumeric string
randx.StringWithCharset(8, "abc")
randx.Hex(4)                // 8-character random hex string

randx.Choice([]int{1,2,3})  // Return a random element

s := []int{1,2,3,4,5}
randx.Shuffle(s)            // Shuffle the slice in place
```

## convert

Data type conversion utilities, supporting safe conversion from any type to common types.

```go
convert.ToString(v)

n, err := convert.ToInt(v)
n, err := convert.ToInt64(v)
f, err := convert.ToFloat64(v)
b, err := convert.ToBool(v)

convert.MustToInt(v)
convert.MustToInt64(v)
convert.MustToFloat64(v)
convert.MustToBool(v)
```

## Testing

```bash
go test ./...
```
