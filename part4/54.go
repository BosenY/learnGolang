package main

import (
	"fmt"
	"math"
)

/*
接口类型是由一组方法来定义的
接口类型的值可以容纳实现这些方法的 any value
这里定义了float类型的MyFloat 的abs  还定义了结构体Vertex的abs 但是不可以使用值类型 必须使用取地址
*/
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser
	// a = v  // a Vertex, does NOT
	// implement Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
