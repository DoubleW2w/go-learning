package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(first(1, 2, 3))
	greet("Hello", "Alice", "Bob", "Charlie")
}

func first(nums ...int) int {
	return nums[0]
}
func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println(prefix, name)
	}
}
