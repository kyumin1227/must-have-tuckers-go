package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	// 한글의 경우 len으로 계산 시 문자 하나 당 3이 나옴
	fmt.Printf("len(str) = %d\n", len(str))		// 12
	fmt.Printf("len(runes) = %d\n", len(runes))	// 8
	fmt.Println(runes)
}