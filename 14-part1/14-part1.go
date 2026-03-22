package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// 数组定义
	var arr1 = [5]int{3, 8, 1, 10, 5}
	max, sum := printArray(arr1[:])
	fmt.Printf("sum:%d, max:%d\n", sum, max)

	// 切片定义
	var slice1 = []int{3, 8, 1, 10, 5}
	max, sum = printArray(slice1)
	fmt.Printf("sum:%d, max:%d\n", sum, max)

	max, sum = printArray([]int{-3, -8, -1, -10, -5})
	fmt.Printf("sum:%d, max:%d\n", sum, max)

	fmt.Println("第 2 题：Slice + Function + Multiple Return Values + Range")
	nums := []int{2, 4, 6, 7, 9, 12}
	evenCount, notEventCount := findEvenAndNotEvent(nums)
	fmt.Printf("偶数有:%d个, 奇数有:%d个\n", evenCount, notEventCount)

	fmt.Println("第 3 题：Map + Range + Function")
	var maps = map[string]int{
		"Alice":   80,
		"Bob":     95,
		"Charlie": 87,
	}
	maxScore, maxName := handleMap(maps)
	fmt.Printf("maxScore:%d, maxName:%s\n", maxScore, maxName)

	fmt.Println("第 4 题：String + Range + Function")
	str := "Go123Lang"
	digitCount, otherCount := countArray(str)
	fmt.Printf("digitCount:%d, otherCount:%d\n", digitCount, otherCount)

}

func printArray(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	max := arr[0]
	var sum int
	for _, v := range arr {
		if v > max {
			max = v
		}
		sum += v
	}
	return max, sum
}

func findEvenAndNotEvent(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	evenCount := 0
	notEventCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		} else {
			notEventCount++
		}
	}
	return evenCount, notEventCount
}

func handleMap(maps map[string]int) (int, string) {
	if maps == nil || len(maps) == 0 {
		return 0, ""
	}
	var maxScore int
	var maxName string
	for k, v := range maps {
		if v > maxScore {
			maxScore = v
			maxName = k
		}
	}
	return maxScore, maxName
}

func countArray(arr string) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	var digitCount int
	var otherCount int
	for _, v := range arr {
		if v >= '0' && v <= '9' {
			digitCount++
		} else {
			otherCount++
		}
	}
	return digitCount, otherCount
}
