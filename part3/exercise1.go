package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i <= 10; i++ {
		z = z - (z*z-x)/2*z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

/*
这个练习让我们实现一个开方函数，使用牛顿法
什么是牛顿法呢？ 这里给自己记录一下：
他是牛顿在17世纪提出的一种在实数域和复数域上近似求解方程的方法
他是通过选择一个初始点z然后重复这一过程求Sqrt(x)的近似值:

z = z - (z * z -x) / 2 * z
这样最后z的值就会不断逼近真正的x的平方根

*/
