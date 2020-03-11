package sampleTest

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) SayHello() string {
	return fmt.Sprintf("hello my name is %s", s.Name)
}

func (s *Student) GoodBye() string {
	return fmt.Sprintf("Bye my name is %s", s.Name)
}
