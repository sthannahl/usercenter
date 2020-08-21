package main

import "fmt"

type A interface {
	Foo()
	Bar()
}

type B struct {
}

func (*B) Foo() {
	fmt.Println(123)
}

func (*B) Bar() {
	fmt.Println(123)
}

type C struct {
	B
}

func (*C) Foo() {
	fmt.Println(456)
}

func main() {
	a := &B{}
	b := C{*a}
	var c A = &b
	c.Foo()
}
