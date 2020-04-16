package main

import "fmt"

type student struct {
	id   int64
	name string
}

//造一个学生管理者
type studentMgr struct {
	allstudent map[int64]student
}

//showStudent
func (s studentMgr) showStudent() {
	for _, stu := range s.allstudent {
		fmt.Printf("学号：%d,姓名：%v\n", stu.id, stu.name)
	}
}

//addStudent
func (s studentMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allstudent[newStu.id] = newStu
	fmt.Println("add success!!!")
}

//editStudent
func (s studentMgr) editStudent() {
	var stuID int64
	fmt.Print("请输入要更改的学号：")
	fmt.Scanln(&stuID)
	value, ok := s.allstudent[stuID]
	if !ok {
		fmt.Println("no this student...")
		return
	}
	fmt.Printf("you editStudent stuID:%d,name:%v\n", value.id, value.name)
	fmt.Print("input student name:")
	var stuName string
	fmt.Scanln(&stuName)
	value.name = stuName
	s.allstudent[stuID] = value
}

//deleteStudent
func (s studentMgr) deleteStudent() {
	var stuID int64
	fmt.Print("input stuID that you want delete: ")
	fmt.Scanln(&stuID)
	_, ok := s.allstudent[stuID]
	if !ok {
		fmt.Println("no this student")
		return
	}
	delete(s.allstudent, stuID)
	fmt.Println("delete success")

}
