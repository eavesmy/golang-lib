package gtype

import (
	"fmt"
	"strconv"
)

// int64 to float64
func Int642Float64(i int64) float64 {
	str := strconv.FormatInt(i, 10)
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

func Float642Int64(i float64) int64 {
	return int64(i)
}

func Int2Float64(i int) float64 {
	return float64(i)
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

func String2Int32(str string) int32 {
	num, _ := strconv.ParseInt(str, 10, 32)
	return int32(num)
}

func String2Float64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

func Float642String(i float64, positions ...int) string {
	position := 6 // 默认6位
	if len(positions) > 0 {
		position = positions[0]
	}

	return fmt.Sprintf("%."+Int2String(position)+"f", i)
}

func Int2String(num int) string {
	return fmt.Sprintf("%d", num)
}

func Int322String(num int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(num)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
