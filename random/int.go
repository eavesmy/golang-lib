package random

import (
	"math/rand"
	"reflect"
	"time"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomIndex(array interface{}) int {
	t := reflect.TypeOf(array).Kind()

	if t != reflect.Slice && t != reflect.Array {
		return -1
	}

	length := reflect.ValueOf(array).Len()

	return RandomInt(0, length-1)
}
