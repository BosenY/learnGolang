package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/*
map就是一个存放键值对的数据结构
map 必须用make来创建；
一个值为nil 的map是空的 并且不能赋值
map的方括号中定义key的类型 后面跟value类型 当然也可以是一个结构体类型
*/

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
