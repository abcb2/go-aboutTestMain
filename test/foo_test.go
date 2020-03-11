package test

import (
	"github.com/abcb2/go-aboutTestMain/foo"
	"testing"
)

func TestFoo_GetName(t *testing.T) {
	f := foo.Foo{Name: "FOO!!"}
	if f.GetName() != "FOO!!" {
		t.Logf("oops!!: %s", f.GetName())
		t.FailNow()
	}
}
