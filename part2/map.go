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

type Z struct {
	x, y int
}

/*
如果顶层类型只有类型名的话，可以将文法的元素中的键名省略
*/
func main() {
	m = make(map[string]Vertex)

	_z := map[string]Z{
		"axx": {1, 2},
		"xxa": {3, 4},
	}
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(_z)
}
