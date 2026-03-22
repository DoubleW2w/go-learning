package main

import "fmt"

func main() {
	// 1. 基本定义
	i := 42
	p := &i // p 指向 i 的地址

	fmt.Printf("i 的值: %d\n", i)
	fmt.Printf("i 的内存地址: %p\n", &i)
	fmt.Printf("指针 p 存储的地址: %p\n", p)
	fmt.Printf("通过指针 p 访问 i 的值: %d\n", *p)

	// 2. 修改值
	*p = 21 // 通过指针修改 i 的值
	fmt.Printf("修改后 i 的值: %d\n", i)

	// 3. 函数传参中的指针
	val := 10
	fmt.Printf("调用前 val: %d\n", val)

	zeroValue(val)
	fmt.Printf("值传递调用后: %d\n", val) // 依然是 10

	zeroPointer(&val)
	fmt.Printf("指针传递调用后: %d\n", val) // 变为 0
}

// 值传递：无法修改原变量
func zeroValue(n int) {
	n = 0
}

// 指针传递：可以修改原变量
func zeroPointer(n *int) {
	*n = 0
}
