package main

import "fmt"

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7 // 校验失败的错误
	ServiceErrCode Code = 8 // 服务错误
)

// 为Code类型添加棒法
func (c Code) GetMsg() string {
	return "成功"
}

// 1. 定义一个基于 int 的自定义类型 MyInt
type MyInt int

// 2. 为 MyInt 添加方法
// IsZero 是一个值接收者方法，它操作的是 MyInt 的副本
func (m MyInt) IsZero() bool {
	return m == 0
}

// Add 是一个指针接收者方法，它可以修改原始值
func (m *MyInt) Add(other MyInt) {
	*m = *m + other
}

// 3. 定义一个结构体，并为其添加方法
type User struct {
	Name string
	Age  int
}

// NewUser 是一个工厂函数，用于创建 User 实例
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// Greet 是 User 的一个值接收者方法
func (u User) Greet() {
	fmt.Printf("你好，我叫 %s，今年 %d 岁。\n", u.Name, u.Age)
}

// CelebrateBirthday 是 User 的一个指针接收者方法，用于修改年龄
func (u *User) CelebrateBirthday() {
	u.Age++
	fmt.Printf("生日快乐！%s 现在 %d 岁了。\n", u.Name, u.Age)
}

func main() {
	fmt.Println(SuccessCode.GetMsg())
	var i int
	fmt.Println(int(SuccessCode) == i) // 必须要转成原始类型才能判断

	fmt.Println("--- 自定义基础类型示例 ---")
	var a MyInt = 10
	fmt.Printf("a 的值: %d, 是否为零: %t\n", a, a.IsZero())

	a.Add(5)
	fmt.Printf("调用 Add(5) 后 a 的值: %d\n", a)

	fmt.Println("\n--- 结构体与方法示例 ---")
	user := NewUser("小明", 18)
	user.Greet()

	user.CelebrateBirthday()
	user.Greet() // 验证年龄是否已更新
}
