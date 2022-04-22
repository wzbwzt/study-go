package splittostring

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//示例
func TestSplittostring1(t *testing.T) { //// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Splittostring("a:b:c", ":") // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

//测试组(无法对单个case进行测试)
func TestSplittostring2(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	test := []testCase{
		{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{str: "adfgadsf", sep: "d", want: []string{"a", "fga", "sf"}},
		{str: "avddfddfddf", sep: "fd", want: []string{"avdd", "d", "df"}},
	}
	for _, v := range test {
		got := Splittostring(v.str, v.sep)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("want:%#v,got:%#v\n", v.want, got)
		}
	}

}

//子测试(即可以组测试，也可以进行单个case的测试)

func TestSplittostring(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testData := map[string]testCase{
		"case 1": {str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case 2": {str: "adfgadsf", sep: "d", want: []string{"a", "fga", "sf"}},
		"case 3": {str: "avddfddfddf", sep: "fd", want: []string{"avdd", "d", "df"}},
	}
	for k, v := range testData {
		t.Run(k, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Splittostring(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("name:%s,want:%#v,got:%#v\n", k, v.want, got)
			}
		})
	}
}

//可以通过-run=RegExp来指定运行的测试用例，还可以通过/来指定要运行的子测试用例，
//例如：go test -v -run=Split/simple只会运行simple对应的子测试用例。

//测试覆盖率：代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是
//在测试中至少被运行一次的代码占总代码的比例。
//使用go test -cover -coverprofile=cover.out  生成覆盖率文件
//使用 go tool cover -html=cover.out    生成html覆盖率的详细文件

//基准测试
//使用go test -bench=Splittostring来查看基准测试结果
//还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据
func BenchmarkSplittostring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Splittostring("adfadfsfsaf", "f")
	}
}

//性能比较测试
//比较同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别；
//性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用
//每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
//可以使用-benchtime标志增加最小基准时间，以产生更准确的结果 go test -bench=Fib40 -benchtime=20s

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

//使用go test -bench=Fib1  来只当执行的测试函数
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//示例函数
func ExampleSplittostring() {
	fmt.Println(Splittostring("adfgadsf", "d"))
	fmt.Println(Splittostring("avddfddfddf", "fd"))
	// Output:
	// [a fga sf]
	// [avdd d df]
}

//goconvey 的使用：
//- go convey是一个支持golang的单元测试框架
//- go convey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到Web界面
//- go convey提供了丰富的断言简化测试用例的编写
func TestIntegerStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func TestSplittostring2convey(t *testing.T) { //// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	Convey("testSplittostring", t, func() {
		//可以使用标准断言，也是可以自己写断言方法
		shouldFunc := func(actual interface{}, expected ...interface{}) string {
			if !reflect.DeepEqual(actual, expected[0]) { // 因为slice不能比较直接，借助反射包中的方法比较
				return fmt.Sprintf("excepted:%v, got:%v", expected, actual) // 测试失败输出错误提示
			}
			return ""
		}
		Convey("test a:b:c |:", func() {
			So(Splittostring("a:b:c", ":"), shouldFunc, []string{"a", "b", "c"})
		})
	})
}

//子测试
func TestSplittostringWithSubTest2convey(t *testing.T) {
	Convey("testSplittostring with subtest", t, func() {
		shouldFunc := func(actual interface{}, expected ...interface{}) string {
			g := actual.([]string)
			w := expected[0].([]string)
			if !reflect.DeepEqual(g, w) { // 因为slice不能比较直接，借助反射包中的方法比较
				return fmt.Sprintf("excepted:%v, got:%v", expected, actual) // 测试失败输出错误提示
			}
			return ""
		}
		Convey("test a:b:c |:", func() {
			So(Splittostring("a:b:c", ":"), shouldFunc, []string{"a", "b", "c"})
		})

		Convey("test adfgadsf | d", func() {
			So(Splittostring("adfgadsf", "d"), shouldFunc, []string{"a", "fga", "sf"})
		})

		Convey("test avddfddfddf | fd", func() {
			So(Splittostring("avddfddfddf", "fd"), shouldFunc, []string{"avdd", "d", "df"})
		})
	})
}

//如果有需要，可以调用 *T 和 *B 的 Skip 方法，跳过该测试或基准测试
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
