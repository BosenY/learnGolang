package main

import "fmt"

/* 当函数的连续多个参数是同一类型,则除了最后一个参数需类型声明之外，其他的都可以去省略 */
func add(x, y int) int {
	return x + y
}

/* 函数可以返回任意数量的返回值 */
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(swap("hello", "word"))
	fmt.Println(split(17))
}
