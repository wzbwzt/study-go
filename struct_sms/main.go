package main

import (
	"fmt"
	"os"
)

//学生管理系统

var stuMgr studentMgr //声明一个全局的学生管理者

//菜单函数
func showMenu() {
	fmt.Println("-------------welcome sms!----------------")
	fmt.Println(`
	1.show all students
	2.add new student
	3.edit student
	4.delete student
	5.exit
	`)
}
func main() {
	stuMgr = studentMgr{
		allstudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		//等待用户输入
		fmt.Print("input action code:")
		var code int8
		fmt.Scanln(&code)
		fmt.Printf("you choice is %v\n", code)
		switch code {
		case 1:
			stuMgr.showStudent()
		case 2:
			stuMgr.addStudent()
		case 3:
			stuMgr.editStudent()
		case 4:
			stuMgr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("error input!!!")
		}
	}

}
