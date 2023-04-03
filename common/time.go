package common

import (
	"strconv"
	"strings"
	"time"
)

const YMDS = "2006-01-02 15:04:05"
const YMD = "2006-01-02"

var zone, _ = time.LoadLocation("Asia/Shanghai")

// 获取当前时间的年月日 时分秒
func GetTimeYMDS() string {
	return time.Now().In(zone).Format(YMDS)
}

func GetTimeByYMD(t int64) int64 {
	t2, _ := time.ParseInLocation(YMD, time.Unix(t, 0).In(zone).Format(YMD), zone)
	return t2.Unix()
}

func GetTodayZero() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}

func TimeToUnix(format string) int64 {
	s := strings.Split(format, ":")
	hour, _ := strconv.Atoi(s[0])
	min, _ := strconv.Atoi(s[1])
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), hour, min, 0, 0, t.Location())
	return newTime.Unix()
}
func DateToNewYorkUnix(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	the_time, err := time.ParseInLocation("2006-01-02", date, loc)
	if err != nil {
		return 0
	}
	return the_time.Unix()
}
func DateTimeToNewYorkUnix(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	if err != nil {
		return 0
	}
	return the_time.Unix()

}
func DateToUnix(date string) int64 {
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02", date, loc)
	if err != nil {
		return 0
	}
	return the_time.Unix()

}
func HourMinuteToUnix(date string) int64 {
	ms := strings.Split(date, ":")
	if len(ms) != 2 {
		return 0
	}
	hour, _ := strconv.Atoi(ms[0])
	minute, _ := strconv.Atoi(ms[1])
	loc, _ := time.LoadLocation("Local")
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, loc).Unix()
}
func GetYearMonthDay() (int, int, int) {
	t := time.Now()
	return t.Year(), int(t.Month()), t.Day()
}
func UnixByYearMonthDay(year, month, day int) int64 {
	t := time.Now()
	newTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}

// 分析收盘状态
func ParseRetreatStatus(startDate string) bool {
	//参数分析
	if startDate == "" {
		return false
	}
	//收盘状态分析
	closingStatus := false
	startClosingTime, _ := time.ParseInLocation("2006-01-02", startDate, zone)
	timeNow := time.Now().Unix()
	if timeNow >= startClosingTime.Unix() {
		closingStatus = true
	}
	return closingStatus
}
