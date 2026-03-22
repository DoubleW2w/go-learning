package main

import (
	"fmt"
)

func main() {
	n := 7
	fmt.Printf("fact(%d) = ", n)
	result := fact(n)
	fmt.Printf(" = %d\n", result)

	// fib 函数递归
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

func fact(n int) int {
	// 有终止条件还不够，还要保证递归的方向
	if n <= 1 {
		fmt.Print("1")
		return 1
	}
	fmt.Printf("%d + ", n)
	return n + fact(n-1)
}
