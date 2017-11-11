package main

import "fmt"

/*
一个结构体就是一个成员变量的集合
结构体的成员变量使用点号来访问，类似于javascript的对象一样
*/

/*
Go 有指针，但没有指针运算
结构体成员变量可以通过结构体指针来访问。
通过指针的间接访问也是透明的
*/

type Classes struct {
	student string
	teacher string
}

/*
结构体有自己的一种文法 structName{}
通过结构体成员变量的值作为列表来新分配一个结构体
使用key: value 预发可以仅列出部分字段(顺序无所谓)

特殊的前缀&构造了指向结构体文法的指针

*/
type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

/*
这里在给自己回忆一下大一学过的C++的指针
int p = 10
这里定义一个整型的变量p 给他赋值10
然后我们去定义一个整形的指针p1
int* p1 = &p
因为 指针就是用来存储地址的变量，所以现在p1就是存储p的存储地址
然后我们使用取变量运算符* 去取p1的变量并且去赋值
*p1 = 20
你会发现 p的值就变成了20
因为p会首先去找它的地址上对应的这个数据，结果发现，它变成了20 所以就是20

当然了 p1也是存储地址的 &p1就知道了它的地址，
但是呢 int* pp = &p1 编译器会出错 为什么呢
因为*p1可以得到地址上int类型的值
但是pp
我们如果那样定义了 结果就是*pp == p1 p1是一个指针类型 不是值类型
所以就知道为什么定义的不对了
我们应该这样来定义 定义一个指向指针的指针
int** pp = &p1
*/

/*
 实际上Go支持只提供类型
*/
func main() {
	var class_24 = Classes{"John", "Nick"}
	fmt.Println(class_24)
	class_24.teacher = "axx"
	fmt.Println(class_24)
	class_23 := &class_24
	class_25 := class_24
	class_25.teacher = "25"
	fmt.Println(class_24) //no difference
	class_23.teacher = "chong"
	fmt.Println(class_24) //show difference
	fmt.Println(class_25) //no difference
	fmt.Println(p, q, r, s)
}
