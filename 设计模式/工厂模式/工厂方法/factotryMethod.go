package main

//工厂方法：
//多个工厂

type ProductType int

const (
	COMPUTER ProductType = iota + 1
	TSHIRT
)

//总工厂
type IProductFactory interface {
	CreateProduct(ProductType) IProduct
}

type IProduct interface {
	GetInfo() string
}

//###########################################
type Computer struct {
}

func (this *Computer) GetInfo() string {
	return "computer"
}

type Tshirt struct {
}

func (this *Tshirt) GetInfo() string {
	return "T-shirt"
}

//###########################################

//技术类工厂
type TechFactory struct{}

func (*TechFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case COMPUTER:
		return &Computer{}
	}
	return nil
}

//日常类工厂
type DailyFactory struct{}

func (*DailyFactory) CreateProduct(t ProductType) IProduct {
	switch t {
	case TSHIRT:
		return &Tshirt{}
	}
	return nil
}

func main() {
	println(new(TechFactory).CreateProduct(COMPUTER).GetInfo())
	println(new(DailyFactory).CreateProduct(TSHIRT).GetInfo())

}
