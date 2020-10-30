package minitools

import (
	"fmt"
	"strconv"
	"time"
)

// TimeSuiteBasic 时间转换基类
type TimeSuiteBasic struct{}

// AnyFormat 时间格式化样式转换
//	"2006/01/02" -> "2006.01.02"
func (ts *TimeSuiteBasic) AnyFormat(oldLayout, newLayout string, t string) string {
	ns, err := time.Parse(oldLayout, t)
	if err != nil {
		panic(err)
	}
	fTime := ns.Format(newLayout)
	return fTime
}

// String2Unix 字符串时间转 Unix 时间戳
//	layout -> 1136142245
func (ts *TimeSuiteBasic) String2Unix(layout string, t string) int64 {
	localTime, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	tString, err := time.ParseInLocation(layout, t, localTime)
	if err != nil {
		panic(err)
	}
	uTime := tString.Unix()
	return uTime
}

// UnixInt2String Unix Int64 转字符串
//	1136142245 -> layout
func (ts *TimeSuiteBasic) UnixInt2String(layout string, t int64) string {
	return time.Unix(t, 0).Format(layout)
}

// Unix2String Unix 转字符串
//	"1136142245" -> layout
func (ts *TimeSuiteBasic) Unix2String(layout string, t string) string {
	if len(t) == 13 {
		t = t[0 : len(t)-3]
	}
	timeInt, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		panic(err)
	}
	sTime := time.Unix(timeInt, 0).Format(layout)
	return sTime
}

// UTC2String UTC 时间转字符串
//	"2006-01-02T03:04:50Z0700" -> layout
func (ts *TimeSuiteBasic) UTC2String(layout string, t string) string {
	rawTime, _ := time.Parse("2006-01-02T15:04:05Z0700", t)
	uTime := rawTime.Format(layout)
	return uTime
}

// RunTime 计算执行时间
func (ts *TimeSuiteBasic) RunTime() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("Time = %v\n", tc)
	}
}
