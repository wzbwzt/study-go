package main



//-------------------
//抽象层
//-------------------
type Car interface {
	Run()
}
type Driver interface {
	Drive(car Car)
}

//-------------------
//实现层
//-------------------
//car
type Benc struct {}
func (b *Benc)Run(){}

type BMW struct {}
func (b *BMW)Run(){}

type Fent struct {}
func (f *Fent)Run(){}

//driver
type ZhanS struct {}
func (f *ZhanS)Drive(car Car){
	car.Run()
}

type LiS struct {}
func (z *LiS)Drive(car Car){
	car.Run()
}

type WanW struct {}
func (f *WanW)Drive(car Car){
	car.Run()
}

//-------------------
//业务逻辑层
//-------------------
func main(){
	// 业务1：张三开奔驰
	var z ZhanS
	var fentian Car
	fentian=&Fent{}  //多态
	z.Drive(fentian)
}