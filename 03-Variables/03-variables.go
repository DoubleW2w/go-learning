package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	f := "string"
	fmt.Println(f)
	var c, d = 1, 2
	fmt.Println(c)
	fmt.Println(d)

	var e int = 3
	fmt.Println(e)

	var h = true
	fmt.Println(h)

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr[0])

	m := map[string]int{
		"张三": 90,
		"李四": 85,
	}
	fmt.Println(m["张三"])
	m["王五"] = 95
	fmt.Println(m)

	g := 10
	p := &g         // 取地址
	fmt.Println(g)  // 10
	fmt.Println(p)  // 地址
	fmt.Println(*p) // 取地址对应的值
}
