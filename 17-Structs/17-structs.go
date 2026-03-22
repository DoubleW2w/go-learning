package main

import "fmt"

func main() {

	p := People{
		Time: "2023-11-15 14:51",
	}
	s := Student{
		People: p,
		Name:   "枫枫",
		Age:    21,
	}
	s.Name = "枫枫知道" // 修改值
	s.PrintInfo()
	s.Info()                   // 可以调用父结构体的方法
	fmt.Println(s.People.Time) // 调用父结构体的属性
	fmt.Println(s.Time)        // 也可以这样

	fmt.Println(s.Age)
	SetAge(s, 18)
	fmt.Println(s.Age)
	SetAge1(&s, 17)
	fmt.Println(s.Age)
}

// Student 定义结构体
type Student struct {
	People
	Name string
	Age  int
}

// PrintInfo 给机构体绑定一个方法
func (s Student) PrintInfo() {
	fmt.Printf("name:%s age:%d\n", s.Name, s.Age)
}

type People struct {
	Time string
}

func (p People) Info() {
	fmt.Println("people ", p.Time)
}

func SetAge(info Student, age int) {
	info.Age = age
}

func SetAge1(info *Student, age int) {
	info.Age = age
}
