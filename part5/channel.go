/*
管道channel 是一种有类型的管道，可以使用channel操作符<-对其发送或者接收值
channel 可以是带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
ch := make(chan int, 100)
*/
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c)
}