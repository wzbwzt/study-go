package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type cat struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

//反射GO中的包为reflect
//反射效率相比正常代码底1~2个数量级

//typeOf
func reflectType(i interface{}) {
	t := reflect.TypeOf(i)
	// fmt.Printf("type:%v\n", t)
	fmt.Printf("type:%v ;kind:%v;%v\n", t.Name(), t.Kind(),t) //Value()有Kind()方法
}

//valueOf
func reflectValue(i interface{}) {
	v := reflect.ValueOf(i)
	k := v.Kind() //TypeOf()有Kind()方法
	fmt.Printf("%v,%v,%T\n", v, k, k)
	switch k {
	case reflect.Float64:
		//v.Float()从反射中获取整型的原始值，然后通过float64()强制转换类型
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	case reflect.Float32:
		//v.Float()从反射中获取整型的原始值，然后通过float32()强制转换类型
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Int64:
		//v.Float()从反射中获取整型的原始值，然后通过int64()强制转换类型
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Int32:
		//v.Float()从反射中获取整型的原始值，然后通过int32()强制转换类型
		fmt.Printf("type is int32,value is %d\n", int32(v.Int()))
	default:
		fmt.Println("other type")
	}

}

//通过反射设置变量的值
func reflectSetValue1(i interface{}) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Float32 {
		value.SetFloat(12.33) //修改的是副本，reflect会引发panic
	}
}
func reflectSetValue2(i interface{}) {
	value := reflect.ValueOf(i)
	if value.Elem().Kind() == reflect.Float32 { //Elem()根据传的指针来获取对应的值；
		value.Elem().SetFloat(12.33)
	}
}

func main() {

	//反射的应用
	//Unmarshal将jsonz字符串给写入为结构体
	str := `{"name":"mimi","age":30}`
	var cat1 cat
	json.Unmarshal([]byte(str), &cat1)
	fmt.Println(cat1)

	//---------------------
	var a int64 = 123
	var f float32 = 3.14
	reflectType(a)
	// reflectValue(a)
	// reflectValue(f)
	var c = cat{}
	// reflectType(c)
	reflectValue(c)
	fmt.Println("============")

	reflectSetValue1(&f)
	fmt.Println(f) //3.14
	reflectSetValue2(&f)
	fmt.Println(f) //12.33

	//isNil()//常被用于判断指针是否为空
	//func (v Value) IsNil() bool
	//报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
	var point *int
	fmt.Println(reflect.ValueOf(point).IsNil())

	//isValid()//常被用于判定返回值是否有效。
	//func (v Value) IsValid() bool
	//返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic
	fmt.Println(reflect.ValueOf(nil).IsValid())
	fmt.Println(reflect.ValueOf("string").IsValid())
	// 实例化一个匿名结构体
	st := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(st).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(st).MethodByName("abc").IsValid())
	ma := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(ma).MapIndex(reflect.ValueOf("娜扎")).IsValid())

}
