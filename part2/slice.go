package main

import "fmt"

/*
go语言不仅有数组，还有切片
go语言切片是对数组的抽象
go数组的长度不可改变，在特定的场合中使用就不太合适，go中提供了一种灵活，功能强大的内置类型切片（动态数组）
与数组相比切片的长度是不鼓捣的，而且可以追加元素
*/
// type Books struct {
// 	title string
// 	author string
// 	subject string
// 	book_id int
//  }
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func main() {
	// var n [10]int /* n 是一个长度为 10 的数组 */

	//定义切片 你可以定义一个未知大小的数组来定义切片
	var new_slice []int
	//也可以使用make()函数来创建一个切片
	// new_slice1 := make([]int, 10, 15) //第二个参数是长度 第三个是最大容器长度，可以不填

	new_slice = append(new_slice, 0)
	printSlice(new_slice)
	new_slice = append(new_slice, 1, 2, 3)
	printSlice(new_slice)

	// 	var i, j int

	// 	/* 为数组 n 初始化元素 */
	// 	for i = 0; i < 10; i++ {
	// 		n[i] = i + 100 /* 设置元素为 i + 100 */
	// 	}

	// 	/* 输出每个数组元素的值 */
	// 	for j = 0; j < 10; j++ {
	// 		fmt.Printf("Element[%d] = %d\n", j, n[j])
	// 	}
}
