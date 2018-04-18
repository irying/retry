package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func func1() {
}

func func2() {
}

type Foo struct {
}

func (f *Foo) func1() {
}

func TestFuncsShouldBeEqual(t *testing.T) {

	foo := &Foo{}
	bar := &Foo{}

	assert.True(t, CompareFuncs(foo.func1, bar.func1),
		"bar.func1 should be equal with foo.func1")
}

func TestFuncsShouldNotBeEqual(t *testing.T) {

	equals := CompareFuncs(func1, func2)
	if equals {
		t.Fatalf("func1 should not be equal with func2")
	}

	foo := &Foo{}
	assert.False(t, CompareFuncs(func1, foo.func1),
		"func1 should not be equal with foo.func1")
}
