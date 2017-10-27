/*
package是定义包名，表明你这个文件属于哪个包, package main 表示一个独立可执行的程序
每个Go应用程序都包含一个名为main的包
*/
package main

import "fmt"
import "./cal"

func main() {
	fmt.Println("Hello, 世界")
	result := cal.Add(1, 2)
	fmt.Println(result)
}
