package main

//命令模式：把请求封装为命令对象，把发出命令的责任和执行命令的责任分割开，可以传递给不同的对象,使得：
//1.命令的发送者和执行者解藕
//2.可以进行命令队列的实现
//3.方便记录执行过程，结合装饰器能更好更方便的扩展

func main() {
	logincmd := NewLoginCommand(&UserSvc{})

	userInvoker := &Invoker{}
	userInvoker.AddCommand(logincmd)
	userInvoker.Do("joel", "123123")
}

//#####################################
type UserSvc struct{}

func (this *UserSvc) Login(name, passw string) error {
	if name != "joel" || passw != "123123" {
		println("login falied")
		return nil
	}
	println("login success")
	return nil
}

type ScoreSvc struct{}

func (this *ScoreSvc) Score() error {
	println("赠送积分成功")
	return nil
}

//####################################
//命令模式
type ICommand interface {
	Exec(args ...interface{}) error
}

//用户登陆命令
type LoginCommand struct {
	*UserSvc
}

func NewLoginCommand(svc *UserSvc) *LoginCommand {
	return &LoginCommand{svc}
}

//TODO 可以添加装饰器，更好的扩展业务
func (this *LoginCommand) Exec(args ...interface{}) error {
	if len(args) != 2 {
		panic("arg panic")
	}
	return this.Login(args[0].(string), args[1].(string))
}

//积分赠送命令
type LoginCommand struct {
	*UserSvc
}

func NewLoginCommand(svc *UserSvc) *LoginCommand {
	return &LoginCommand{svc}
}

//TODO 可以添加装饰器，更好的扩展业务
func (this *LoginCommand) Exec(args ...interface{}) error {
	if len(args) != 2 {
		panic("arg panic")
	}
	return this.Login(args[0].(string), args[1].(string))
}

//执行者:可以做队列等
type Invoker struct {
	cmds []ICommand
}

func (this *Invoker) Do(args ...interface{}) error {
	for _, v := range this.cmds {
		v.Exec(args...)

	}
	return nil
}

func (this *Invoker) AddCommand(cmds ...ICommand) {
	this.cmds = append(this.cmds, cmds...)
}
