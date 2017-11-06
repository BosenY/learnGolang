package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
//使用一个闭包函数来实现最经典的斐波那契函数
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		res := sum1 + sum2
		sum1 = sum2
		sum2 = res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
