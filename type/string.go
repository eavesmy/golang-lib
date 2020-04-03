package gtype

// Reverser string
// str := Reverse("123") // "321"
func Reverse(str string) string {
	b := []rune(str)

	length := len(b)
	r := make([]rune, length, length)

	for i, item := range b {
		r[length-1-i] = item
	}

	return string(r)
}