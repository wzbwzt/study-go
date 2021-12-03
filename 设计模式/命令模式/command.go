package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

//命令模式：把请求封装为命令对象，把发出命令的责任和执行命令的责任分割开，可以传递给不同的对象,使得：
//1.命令的发送者和执行者解藕
//2.可以进行命令队列的实现
//3.方便记录执行过程，结合装饰器能更好更方便的扩展

//#####################################
type UserSvc struct{}

func (this *UserSvc) Login(name, passw string) error {
	if name != "joel" || passw != "123123" {
		println("login falied")
		return errors.New("账号/密码错误")
	}
	println("login success")
	return nil
}

type ScoreSvc struct{}

func (this *ScoreSvc) Score(name string) error {
	println(name + "赠送积分成功")
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
type ScoreCommand struct {
	*ScoreSvc
}

func NewScoreCommand(svc *ScoreSvc) *ScoreCommand {
	return &ScoreCommand{svc}
}

//TODO 可以添加装饰器，更好的扩展业务
func (this *ScoreCommand) Exec(args ...interface{}) error {
	if len(args) < 1 {
		panic("arg panic")
	}
	return this.Score(args[0].(string))
}

//执行者:可以做队列等
type Invoker struct {
	cmds []ICommand
}

func (this *Invoker) Do(args ...interface{}) error {
	for _, v := range this.cmds {
		err := v.Exec(args...)
		if err != nil {
			break
		}

	}
	return nil
}

func (this *Invoker) AddCommand(cmds ...ICommand) {
	this.cmds = append(this.cmds, cmds...)
}

//############################################
//有预谋
var UserInvoker *Invoker

func init() {
	logincmd := NewLoginCommand(&UserSvc{})
	scorecmd := NewScoreCommand(&ScoreSvc{})

	UserInvoker = &Invoker{}
	UserInvoker.AddCommand(logincmd, scorecmd)
	// userInvoker.Do("joel", "123123")
}

func main() {
	e := gin.Default()
	e.POST("/users", func(c *gin.Context) {
		type User struct {
			Name     string `json:"name"`
			PassWord string `json:"passWord"`
		}

		param := &User{}
		c.ShouldBind(param)

		UserInvoker.Do(param.Name, param.PassWord)

	})

	e.Run(":8088")
}
