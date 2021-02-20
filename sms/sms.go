package sms

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var numeric = [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func GenCode(widths ...int) string {
	width := 4
	if len(widths) > 0 {
		width = widths[0]
	}

	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
