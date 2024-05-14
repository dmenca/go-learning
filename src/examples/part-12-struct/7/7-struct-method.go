package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 给Student类型定义方法
func (s *Student) PrintInfo() {
	fmt.Printf("Name:%s,Age:%d\n", s.Name, s.Age)
}

func (s *Student) AddAge(age int) {
	s.Age = s.Age + age
}

func main() {

	student := Student{Name: "张三", Age: 20}
	// 调用Student类型的方法
	student.PrintInfo()

	// 调用Student类型的方法
	student.AddAge(5)
	student.PrintInfo()
}
