package main

import "fmt"

func main() {
	var char rune = '한'	// rune 타입은 int32 타입과 동일

	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
}