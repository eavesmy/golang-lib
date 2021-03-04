package date

import "time"

const (
	YMDhms = "2006-01-02 15:04:05"
)

// 获取当天 00:00:00 - 23:59:59
func DayRange() (s, e string) {

	d := time.Now()

	start := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	end := time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 59, d.Location())

	s = start.Format(YMDhms)
	e = end.Format(YMDhms)

	return
}

func TodayFormat(format string) string {
	return time.Now().Format(format)
}
