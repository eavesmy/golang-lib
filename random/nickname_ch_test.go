package random

import "testing"

func NickName_ChTest(t *testing.T) {
	name := NickName_Ch()
	t.Fatal(name)
}
