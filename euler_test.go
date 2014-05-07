package euler

import (
	"io"
	"os"
	"reflect"
	"runtime"
	"testing"
)

type problem func() string
type problemWithInput func(io.Reader) string

func (p problem) String() string {
	n := runtime.FuncForPC(reflect.ValueOf(p).Pointer()).Name()
	return n[len(n)-3:]
}

func (p problemWithInput) String() string {
	n := runtime.FuncForPC(reflect.ValueOf(p).Pointer()).Name()
	return n[len(n)-3:]
}

func test(t *testing.T, p problem, s string) {
	a := p()
	if a != s {
		t.Errorf("Problem %s != %v\n", p, a)
	}
}

func testWithInput(t *testing.T, p problemWithInput, s string) {
	in, err := os.Open("testdata/" + p.String() + ".txt")
	if err != nil {
		t.Errorf("Problem %s needs testdata to be run.", p)
	}

	a := p(in)
	if a != s {
		t.Errorf("Problem %s != %v\n", p, a)
	}
}
