package main

import "fmt"
type A struct{}
func (a A) Foo() { fmt.Println("A.Foo") }

type B struct{}
func (b B) Foo() { fmt.Println("B.sssBar") }

type C struct {
	A
	B
}


func main(){
	var s= C{}
	s.A.Foo()
}