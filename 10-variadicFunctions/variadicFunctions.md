一个函数只能有一个可变参数，并且它必须是最后一个参数。

```go
func sum(nums []int) {} → sum([]int{1, 2, 3})
func sum(nums ...int) {} → sum(1, 2, 3)
```

```go
nums := []int{1, 2, 3}
func sum(values ...int) {}
// 应该如何调用 sum？
```
在 Go 里，如果一个函数参数是 ...int，而你手里有一个 []int，调用时需要用一个特殊写法，把切片“展开”为多个参数。
```go
sum(nums...)
```

数组 array
- 定义： 长度固定的一组同类型元素。

```go
[3]int
// 元素类型是 int
// 长度固定是 3
```

切片 slice
- 定义： 对一段同类型元素的“动态视图”。

```go
[]int
// 元素类型是 int
// 长度是动态的
// 容量是底层数组的长度减去切片的起始索引
```

可变参数 variadic parameter
- 定义： 函数参数的一种特殊写法，表示“可以传入任意多个同类型参数”。

```go
...int
```
```go
func sum(nums ...int)
```
只能用于函数定义
必须是最后一个参数
一个函数里通常只能有一个可变参数
在函数内部，它表现得像一个 []int

---

```go
// 切片调用
func f(s []int)
f([]int{1,2,3})

// 可变参数
func f(nums ...int)
f(1,2,3)

// 可变参数函数，切片参数（解构）
nums := []int{1,2,3}
func f(values ...int)
f(nums...)

```
