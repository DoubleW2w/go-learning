package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// 生成一组递增常量的工具
const (
	A = iota
	B
	C
	_ // 跳过某些值
	E
)

func main() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	fmt.Println(A, B, C, E)

}
