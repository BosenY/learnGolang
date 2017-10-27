package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	if a := 2; a < 10 {
		fmt.Println("a is belowed than 10")
	}
}
