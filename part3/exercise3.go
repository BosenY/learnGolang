package main

import "fmt"

//立方根
func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for i := 0; i <= 10; i++ {
		z = z - (z*z*z-x)/(3*z*z)
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
