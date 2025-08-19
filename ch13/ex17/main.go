package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hello"
	addr1 := unsafe.StringData(str)
	str += " World"
	addr2 := unsafe.StringData(str)
	str += " Welcome!"
	addr3 := unsafe.StringData(str)

	// 문자열은 합 연산 수행 시 새로운 문자열로 반환
	fmt.Println(str)
	fmt.Printf("addr1:\t%p\n", addr1)
	fmt.Printf("addr2:\t%p\n", addr2)
	fmt.Printf("addr3:\t%p\n", addr3)
}