package foo

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("foo main_test start")
	code := m.Run()
	println("foo main_test fin")
	os.Exit(code)
}
