package main

import "fmt"

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

// 断言类型
// func sing(obj Animal) {
//   // 通过断言来获取此时的具体类型
//   switch obj.(type) {
//   case Chicken:
//     fmt.Println("鸡")
//   case Cat:
//     fmt.Println("猫")
//   }
//   obj.sing()
// }

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// 全部实现完之后，chicken就不再是一只普通的鸡了
func main() {
	var animal Animal

	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()

}
