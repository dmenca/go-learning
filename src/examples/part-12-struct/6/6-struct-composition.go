package main

import (
	"fmt"
)

type Class struct {
	Name string
}
type Student struct {
	Name  string
	class *Class
}

type Student2 struct {
	Name  string
	class Class
}

func main() {
	c := Class{Name: "二班"}
	student := Student{Name: "张三", class: &c}
	fmt.Println(student.class.Name)
	student2 := Student2{Name: "张三", class: c}
	fmt.Println(student2.class.Name)
	c.Name = "三班"
	// Student中的class是指针类型，因此c的内容改变了，student中的数据也会改变
	fmt.Printf("student1 class name:%s, address: %p\n", student.class.Name, &student.class)
	// Student中的class不是指针类型，是正常的变量声明，相当于拷贝了一个变量进去，外部的变量改变，内部的变量不会改变
	fmt.Printf("student2 class name:%s, address: %p\n", student2.class.Name, &(student2.class))

	student2.class.Name = "四班"

	fmt.Printf("student1 class name:%s, address: %p\n", student.class.Name, &student.class)
	fmt.Printf("student2 class name:%s, address: %p\n", student2.class.Name, &student.class)

}
