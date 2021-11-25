package main

type user struct {
	name string
}

func (user) print1() {
	println("p1")
}

func (*user) print2() {
	println("p1")
}
func main() {
	var u user
	user.print1(u)

	var up *user
	(*user).print2(up)

	// (*user).print1(up)
	// user.print2(u)

}
