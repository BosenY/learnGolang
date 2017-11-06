package main

import (
	"fmt"
	"math"
)
/*
不仅仅包括结构体，其他的类型也可以定义方法
 */
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
