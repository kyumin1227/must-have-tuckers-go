package main

import "fmt"

func Multiple(a, b int) int {
	return a * b
}

func main() {
	result := Multiple(3, 6)
	fmt.Println(result)
}