package foo

import "testing"

func TestFoo_GetName(t *testing.T) {
	foo := &Foo{Name: "hogehoge"}
	if foo.GetName() != "hogehoge" {
		t.Error("oops")
	}
}
