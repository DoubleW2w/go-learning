package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	res := add(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res1, res2 := sub(1, 2)
	fmt.Println("1-2 =", res1, "2-1 =", res2)
}

func add(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func sub(a, b int) (int, int) {
	return a - b, b - a
}
