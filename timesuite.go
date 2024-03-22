package minitools

import (
	"fmt"
	"strconv"
	"time"
)

// TimeSuiteBasic
type TimeSuiteBasic struct{}

// AnyFormat any format to any
//
//	"2006/01/02" -> "2006.01.02"
func (ts *TimeSuiteBasic) AnyFormat(oldLayout, newLayout string, t string) (string, error) {
	ns, err := time.Parse(oldLayout, t)
	if err != nil {
		return "", err
	}
	fTime := ns.Format(newLayout)
	return fTime, nil
}

// String2Unix String(layout) time to Unix
//
//	layout -> 1136142245
func (ts *TimeSuiteBasic) String2Unix(layout string, t string) (int64, error) {
	localTime, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	tString, err := time.ParseInLocation(layout, t, localTime)
	if err != nil {
		return 0, err
	}
	uTime := tString.Unix()
	return uTime, nil
}

// UnixInt2String Unix(int) to String(layout)
//
//	1136142245 -> layout
func (ts *TimeSuiteBasic) UnixInt2String(layout string, t int64) string {
	return time.Unix(t, 0).Format(layout)
}

// Unix2String Unix(string) to String(layout)
//
//	"1136142245" -> layout
func (ts *TimeSuiteBasic) Unix2String(layout string, t string) (string, error) {
	if len(t) == 13 {
		t = t[0 : len(t)-3]
	}
	timeInt, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return "", err
	}
	sTime := time.Unix(timeInt, 0).Format(layout)
	return sTime, nil
}

// UTC2String UTC to String(layout)
//
//	"2006-01-02T03:04:50Z0700" -> layout
func (ts *TimeSuiteBasic) UTC2String(layout string, t string) (string, error) {
	rawTime, err := time.Parse("2006-01-02T15:04:05Z0700", t)
	if err != nil {
		return "", err
	}
	uTime := rawTime.Format(layout)
	return uTime, nil
}

// RunTime
func (ts *TimeSuiteBasic) RunTime() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("Time = %v\n", tc)
	}
}
