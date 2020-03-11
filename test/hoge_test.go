package test

import (
	sampleTest "github.com/abcb2/go-aboutTestMain"
	"testing"
)

func TestStudent(t *testing.T) {
	student := &sampleTest.Student{
		Name: "Taro",
		Age:  20,
	}
	t.Run("SayHello", func(t *testing.T) {
		if student.SayHello() != "hello my name is Taro" {
			t.FailNow()
		}
	})
	t.Run("Bye", func(t *testing.T) {
		if student.GoodBye() != "Bye my name is Taro" {
			t.Logf("oops '%s'", student.GoodBye())
			t.FailNow()
		}
	})
}
