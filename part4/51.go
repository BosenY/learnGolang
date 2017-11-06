package main

import (
	"fmt"
	"math"
)

/*
在Go中 没有类的概念
但是你可以在结构体上定义方法
方法接收者出现在func关键字和方法名之间的参数中
*/
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { //在Vertex上定义了方法Abs
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	s := &Vertex{3, 4}
	fmt.Println(s.Abs())
}
