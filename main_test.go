package sampleTest

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("----------START")
	code := m.Run()
	println("----------FINISH")
	os.Exit(code)
}
