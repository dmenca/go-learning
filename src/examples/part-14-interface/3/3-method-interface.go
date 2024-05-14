package main

import (
	"fmt"
	"math"
)

// 定义一个接口 接口中定义了Area方法 等于被实现
type Shape interface {
	Area() float64
}

// 定义结构体Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现Shape的Area接口
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 6}
	shapes := []Shape{&circle, &rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}
}
