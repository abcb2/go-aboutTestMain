package bar

import "testing"

func Test_myName(t *testing.T) {
	if myName != "BARBAR" {
		t.Error("oops2")
	}
}
