package foo

type Foo struct {
	Name string
}

func (foo *Foo) GetName() string {
	return foo.Name
}
