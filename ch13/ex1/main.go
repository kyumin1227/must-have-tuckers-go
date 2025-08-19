package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n"
	// 백틱으로 묶을 경우 특수 문자가 동작하지 않음
	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}