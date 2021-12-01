package main

//工厂方法：
//多个工厂

//总工厂
type IProductFactory interface {
	CreateProduct()
}

type IProduct interface {
	GetInfo() string
}

//技术类工厂
type TechFactory struct{}

func (*TechFactory) CreateProduct() IProduct {
	return &TechFactory{}
}

func (*TechFactory) GetInfo() string {
	return "computer"
}

//日常类工厂
type DailyFactory struct{}

func (*DailyFactory) CreateProduct() IProduct {
	return &DailyFactory{}
}

func (*DailyFactory) GetInfo() string {
	return "T-shirt"
}

func main() {
	println(new(TechFactory).CreateProduct().GetInfo())
	println(new(DailyFactory).CreateProduct().GetInfo())

}
