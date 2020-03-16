package bar

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("======>bar_START")
	code := m.Run()
	println("======>bar_FIN")
	os.Exit(code)
}
