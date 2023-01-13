# 1.18 新增功能之泛型

类型参数
Go 语言中的“泛型”是通过支持类型参数实现的，类型参数又可以分为“函数的类型参数”，“类型的类型参数”和“方法的类型参数”。

## 函数的类型参数 - 泛型函数

```go
func MinAny[T int](x, y T) T {
    if x < y {
        return x
    }
    return y
}
```

- 在函数名和参数列表之间的 [T int]，这就是类型参数，我们在函数 MinAny 中使用类型参数，该函数就是“泛型函数”;
- 类型参数支持多个类型，使用 | 分隔，例如：[T int | float64];
- 当类型参数支持的数值类型很多时可以声明一个接口类型，不同的是，接口类型中不再是函数，而是类型
- 如果 [] 中包含多个类型参数，需要使用英文逗号 , 分隔，并且类型参数的形参名字不能相同，例如：[T ordered, T1 ordered1

```go
//波浪线开头代表类型本身和以该类型为底层类型的所有类型
type ordered interface {
    ~int | ~float64
}

func MinAny[T ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}
```

## 类型的类型参数 - 泛型类型

```go
type MinSalary[T int | float64] struct {
    salary T
}
```

- 定义一个自定义类型 MinSalary，它是一个“泛型类型”，与定义一个自定义“普通类型”的区别是在类型名字后面跟一个[]中括号，里面包含类型参数（其中 T 是类型形参，int 和 float64 是类型实参）

- “泛型类型”和“泛型函数”使用方式不同，它不能像“泛型函数”具备类型推断的功能，而是需要显示指定类型实参

```go
salary := &MinSalary[int]{
    salary: 1000,
}
fmt.Printf("%+v\n", salary)

```

## 方法的类型参数 - 泛型方法

```go
type Salary[T int | float64] struct {
    money T
}

func (s *Salary[T]) Min(x, y T) T {
    if x < y {
        return x
    }
    return y
}
```

- 方法的接收者除了类型名称之外，还有类型参数的形参 `*Salary[T]`
- “泛型方法”不能像“泛型函数”那样，具有自身的类型参数，以下代码目前是不支持的。

```go
func (s *Salary[T]) Min[T1 int](x, y T) T {
    if x < y {
        return x
    }
    return y
}
```
