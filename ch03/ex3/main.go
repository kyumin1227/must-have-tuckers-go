package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n", a)
	fmt.Printf("%08.2g\n", a)
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%f\n", c)
	// 00324.13
	// 03.2e+02
	// 324.13
	// 3.140000
}