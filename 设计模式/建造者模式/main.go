package main

import "fmt"

//建造者模式，规划实例的构建，和参数的规范
type Book struct {
	ID    int64   //必须传
	Name  string  //必须传
	Price float32 //不可以为负数
}

func (this *Book) Builder(id int64, name string) *bookBuilder {
	return &bookBuilder{id: id, name: name}
}

//==============================================================================
//创建构造体，可以随着升级而改变
type bookBuilder struct {
	id    int64
	name  string
	price float32
}

func (this *bookBuilder) SetPrice(price float32) *bookBuilder {
	if price > 0 {
		this.price = price
	}
	return this
}
func (this *bookBuilder) Build() *Book {
	return &Book{ID: this.id, Name: this.name, Price: this.price}
}

//==============================================================================
func main() {
	book := new(Book).Builder(123, "joel").SetPrice(112.3).Build()
	fmt.Printf("%v", book)
}
