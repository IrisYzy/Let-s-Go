package main

import "fmt"

func f(p *int) int {
	*p = 123
	return *p
}

func foo() int {
	var x int
	y, _ := x, f(&x)
	return y
}

func bar() int {
	var x int
	var y, _ = x, f(&x)
	return y
}

func testGo()  {
	fmt.Println("22222")
}

func main() {

	println(foo(), bar())
}