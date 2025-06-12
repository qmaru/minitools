package time

import (
	"fmt"
	"strconv"
	"time"
)

// TimeSuiteBasic
type TimeSuiteBasic struct{}

// ConvertFormat convert format to any
//
//	"2006/01/02" -> "2006.01.02"
func (ts *TimeSuiteBasic) ConvertFormat(oldLayout, newLayout, t string) (string, error) {
	parsed, err := time.Parse(oldLayout, t)
	if err != nil {
		return "", fmt.Errorf("failed to parse time '%s' with layout '%s': %w", t, oldLayout, err)
	}
	return parsed.Format(newLayout), nil
}

// ParseTimeToUnix String(layout) time to Unix
//
//	layout -> 1136142245
func (ts *TimeSuiteBasic) ParseTimeToUnix(layout, t string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, fmt.Errorf("load location failed: %w", err)
	}

	parsed, err := time.ParseInLocation(layout, t, loc)
	if err != nil {
		return 0, fmt.Errorf("failed to parse time '%s': %w", t, err)
	}

	return parsed.Unix(), nil
}

// FormatUnixToString Unix(int) to String(layout)
//
//	1136142245 -> layout
func (ts *TimeSuiteBasic) FormatUnixToString(layout string, t int64) string {
	return time.Unix(t, 0).Format(layout)
}

// ParseUnixStringToTime Unix(string) to String(layout)
//
//	"1136142245" -> layout
func (ts *TimeSuiteBasic) ParseUnixStringToTime(layout, t string) (string, error) {
	if len(t) == 13 {
		t = t[0 : len(t)-3]
	}

	seconds, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid unix string '%s': %w", t, err)
	}

	return time.Unix(seconds, 0).Format(layout), nil
}

// UTCStringToLocal UTC to String(layout)
//
//	ISO8601/RFC3339 (2006-01-02T03:04:50Z0700) -> layout
func (ts *TimeSuiteBasic) UTCStringToLocal(layout string, t string) (string, error) {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return "", fmt.Errorf("failed to parse UTC time '%s': %w", t, err)
	}
	return parsed.Format(layout), nil
}

// ElapsedTime calc time
func (ts *TimeSuiteBasic) ElapsedTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time = %v\n", time.Since(start))
	}
}

func New() *TimeSuiteBasic {
	return new(TimeSuiteBasic)
}
