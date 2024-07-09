package test

import (
	"testing"
)

type a struct {
	name string
}

func TestMytest(t *testing.T) {
	b := &a{}
	c := &a{}

	b.name = "b"
	c.name = "a"

	t.Logf("b : %s, c : %s", b.name, c.name)
}
