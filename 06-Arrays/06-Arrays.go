package main

import "fmt"

func main() {
	fmt.Println("start")
	// 声明一个空数组
	var arr [5]int
	fmt.Println("len is ", len(arr))
	fmt.Println("arr is ", arr)

	// 声明一个空数组
	var arrb = [5]string{"Go", "Java", "Python"}
	fmt.Printf("1:%s ,3:%s \n", arrb[0], arrb[2])
	arrb[2] = "aaaa"
	fmt.Printf("1:%s ,3:%s \n", arrb[0], arrb[2])

	arrC := [5]string{"Go", "Java", "Python"}
	fmt.Println(arrC)
}
