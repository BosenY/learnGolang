/*
package是定义包名，表明你这个文件属于哪个包, package main 表示一个独立可执行的程序
每个Go应用程序都包含一个名为main的包
*/
package main 

/* 
这里import就是去导入你所要使用的包，这里导入了一个用于打印控制台的fmt包和一个自己写的cal包 
*/
import "fmt"
import "./cal"

/*
func main() 是一个函数 而且main()是一个程序开始执行的函数，main函数式每一个可执行程序必须包含的
*/
func main() {
	fmt.Println("Hello, 世界")
	result := cal.Add(1, 2)
	fmt.Println(result)
}
