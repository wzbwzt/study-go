package main

import "fmt"

func main() {
	t := MinAny(2, 1)
	fmt.Println(t)

	salary := MinSalary[float32]{salary: 1000}
	fmt.Println(salary)

	currentSalary := salary.upSalary(1000.88)
	println(currentSalary)

}

type orderNum interface {
	~int | ~float32 | ~int64
}

//泛型函数
func MinAny[T orderNum](x, y T) T {
	if x < y {
		return x
	}
	return y
}

//泛型类型
type MinSalary[T int | float32] struct {
	salary T
}

//泛型方法
func (m *MinSalary[T]) upSalary(x T) T {
	m.salary += x
	return m.salary
}
