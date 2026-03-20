package main

import (
	"fmt"
	"time"
)

func main() {
	// 简化声明在循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 10 {
		fmt.Println("range", i)
		if i == 5 {
			fmt.Println("current ", i)
		}
		if i%2 == 0 {
			fmt.Println("even")
		}
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	cc := 2
	fmt.Print("Write ", cc, " as ")
	switch cc {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
