package main

import (
	"fmt"
)

func test(name string) (w string) {
	/*	./return.go:8: w redeclared in this block
		previous declaration at ./return.go:7*/
	var w string
	w = name + " world"
	return w
}

func test1(name string) string {
	var w string
	w = name + " world"
	return w
}

func main() {
	var n string
	name := "hello"
	n = test1(name)
	fmt.Println(n)
}
