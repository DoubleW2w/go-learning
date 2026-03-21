package main

import (
	"fmt"
)

// intSeq 函数返回一个在其函数体内定义的匿名函数。
// 返回的函数使用闭包的方式 隐藏 变量 i，
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// intSeq 函数，将返回值（一个函数）赋给 nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	// 这是一个切片，切片元素是函数类型，
	// func() int 表示 无参数，返回int
	funcs := []func() int{}

	for i := 0; i < 3; i++ {
		// 将一个匿名函数放进切片中，（这个匿名函数是返回i的函数）
		funcs = append(funcs, func() int { return i })
	}

	// 打印 funcs[0](), funcs[1](), funcs[2]()
	fmt.Println(funcs[0]()) // 0
	fmt.Println(funcs[1]()) // 1
	fmt.Println(funcs[2]()) // 2

}
