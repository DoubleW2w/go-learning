# Go 语言自定义类型与方法

在 Go 中，你可以使用 `type` 关键字基于现有类型创建新的、独特的类型。这不仅仅是别名，而是创建了一个全新的类型，可以为其附加特定的方法。

## 1. 定义自定义类型

语法：`type NewTypeName BaseType`

```go
// 定义一个名为 MyInt 的新类型，其底层是 int
type MyInt int
```

- `MyInt` 是一个全新的类型，它和 `int` 是不同的。
- 变量 `var a int` 和 `var b MyInt` 不能直接相互赋值，需要显式转换：`a = int(b)`。

```go
// 案例2
type Age int
var a int = 3
var age Age = Age(a)  // 类型转换：int → Age
```

## 2. 类型转换
在 Go 里，string(某个整数) 更像是——拿这个整数去字符表里找对应的字符，而不是把数字本身“拼成文本”。
- `string(65)` 的结果不是数字文本 "65"
- 而是一个只包含一个字符的字符串，这个字符是编号 65 对应的字符，也就是 "A"
是按照目标类型的规则来解释这个值

## 3. 为自定义类型添加方法

方法是附加到**特定类型**上的函数。通过方法，我们可以为自定义类型赋予行为。

语法：`func (receiver ReceiverType) MethodName(parameters) (return_types)`

- **Receiver**：它将函数与类型关联起来。可以是值接收者或指针接收者。
- 在 Go 里，像 int 这种内建类型本身，你不能直接随意给它加上你自己的业务方法
- 而 `type Age int` 这种你自己定义的命名类型，才是方法接收者的常见载体
```go
// 为 MyInt 类型添加一个 IsPositive 方法
func (m MyInt) IsPositive() bool {
    return m > 0
}
```

### 值接收者 vs. 指针接收者

- **值接收者 (`func (t T)`)**：
  - 方法操作的是类型的**副本**。
  - 不会修改原始值。
  - 适用于不需要修改状态的场景。

- **指针接收者 (`func (t *T)`)**：
  - 方法操作的是类型的**原始值**（通过指针）。
  - 可以修改原始值。
  - 适用于需要修改状态或处理大型结构体的场景，以避免复制开销。

所以这里判断**值接收者**还是**指针接收者**，关键不在“副本会不会不准确”，而在于：
- 这个方法是否要修改数据？
- 这个类型是否大到值得避免复制？
- 语义上，它更像“读一个值”，还是“操作一个对象”？


## 类型别名
和自定义类型很像，但是有一些地方和自定义类型有很大差异
1. 不能绑定方法
2. 打印类型还是原始类型
3. 和原始类型比较，类型别名不用转换
```go
package main

import "fmt"

type AliasCode = int
type MyCode int

const (
  SuccessCode      MyCode    = 0
  SuccessAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
func (m AliasCode) MyAliasCodeMethod() {

}

func main() {
  // 类型别名，打印它的类型还是原始类型
  fmt.Printf("%T %T \n", SuccessCode, SuccessAliasCode) // main.MyCode int
  // 可以直接和原始类型比较
  var i int
  fmt.Println(SuccessAliasCode == i)
  fmt.Println(int(SuccessCode) == i) // 必须转换之后才能和原始类型比较
}
```


## 3. 为什么使用自定义类型？

- **增强代码可读性**：`type UserID string` 比 `string` 更能表达意图。
- **类型安全**：防止将不同业务含义但底层类型相同的数据混用。如果都用 int，编译器无法帮你区分“这个整数到底表示年龄、分数还是价格”
这样就容易发生参数传错、字段混用、业务含义混乱的问题，即就是让类型系统参与业务约束
- **封装行为**：将数据和操作该数据的方法绑定在一起，实现面向对象的编程思想。
