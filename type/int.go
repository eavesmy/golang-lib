package gtype

import "strconv"

// int64 to float64
func Int642Float64(i int64) float64 {
	str := strconv.FormatInt(i, 10)
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

func Float642Int64(i float64) int64 {
	return int64(i)
}

func Int2Int64(i int) int64 {
	return int64(i)
}

func Int642Int(i int64) int {
	str := strconv.FormatInt(i, 10)
	num, _ := strconv.Atoi(str)
	return num
}

func Int642String(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

func String2Int64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func String2Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
