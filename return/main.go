package main

import "fmt"

/*
return 会做什么呢?

Go 的函数返回值是通过堆栈返回的, return 语句不是原子操作，而是被拆成了两步.

* 给返回值赋值 (rval)
* 调用 defer 表达式
* 返回给调用函数(ret)

*/
func main() {
	test1(1)
	test2()
	test3()
	test4()
	test5()

}

func test1(d int) (ret int) {
	defer func() {
		ret++
	}()

	return d

	//output:
	//2
}

func test2() (result int) {
	defer func() {
		result++
	}()
	return 0

	/*
		output:
		1
	*/
}

func test3() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t

	/*
		output:
		5
	*/
}

func test4() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
	/*
		outout:
		1
		//defer函数中传的是r的复制值，只是名称也是r,不再是返回的r
	*/
}

type Test struct {
	Max int
}

func (t *Test) Println() {
	fmt.Println(t.Max)
}

func deferExec(f func()) {
	f()
}

func test5() {
	var t *Test
	defer deferExec(t.Println)

	t = new(Test)
	/*
		output:
		panic
		//defer函数中传入t时还是nil,nil调用Println会panic
	*/
}
