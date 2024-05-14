package main

import "fmt"

type VideoCourse struct {
	Name        string
	Language    string
	TimeMinutes int
}

type Student struct {
	Name string
	Age  int
}

func main() {
	course := VideoCourse{}
	course.Name = "Go Language Overview"
	course.Language = "Go"
	course.TimeMinutes = 60

	fmt.Println("Video course:", course)

	student := Student{Name: "张三", Age: 16}
	fmt.Println("Student:", student)

}
