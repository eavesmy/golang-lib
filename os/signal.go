package os

// Usage:
/*
import gos "github.com/eavesmy/golang-lib/os"

func main(){
	gos.N().Handle(syscall.SIGCLD,func(){
		// handle
	}).L()
}

*/

import (
	"os"
	"os/signal"
)

type sig struct {
	c chan os.Signal
	m map[os.Signal]func()
}

func (s *sig) Handle(code os.Signal, handle func()) *sig {
	signal.Notify(s.c, code)
	s.m[code] = handle
	return s
}

func (s *sig) L() {
	for {
		c := <-s.c
		if h, exists := s.m[c]; exists {
			h()
		}
	}
}

func N() (s *sig) {
	s = &sig{c: make(chan os.Signal), m: map[os.Signal]func(){}}
	return
}
