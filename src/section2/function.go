package main

import (
	"fmt"
)

func main() {

	a := add(2, 3)
	fmt.Println("a =", a)

	f := func(a int, b int) int { return a + b }
	fmt.Println(f(1, 2))

	fmt.Println(fuga(20, 10))

	fmt.Println(div(20, 10))
}

func add(a int, b int) int {
	return a + b
}

func fuga(a int, b int) (sum int) {
	sum = a + b
	return
}

func div(a int, b int) (int, int) {
	var remainder = a % b
	return a / b, remainder
}
