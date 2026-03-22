package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")

	// func->defer2->defer1->defer4->defer3
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
}

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Func() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}
