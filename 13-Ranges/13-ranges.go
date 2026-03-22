package main

import "fmt"

func main() {
	// 遍历一个数组
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}

	for _, v := range arr {
		fmt.Printf(" = %d\n", v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// 中文占三个字节，i是起始字节位置，而且索引从0开始
	// 你 占了 0,1,2 ，所以好从3开始
	// rune 对应的是字符的码点值即unicode编码值，表示是一个字符的“身份证”信息
	for i, c := range "你好" {
		fmt.Printf("%d, %d, %c,\n", i, c, c)
	}
}
