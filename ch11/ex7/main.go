package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A	int8	// 1
	B	int		// 8
	C	int8	// 1
	D	int		// 8
	E	int8	// 1
}

func main() {
	user := User{ 1, 2, 3, 4, 5 }
	fmt.Println(unsafe.Sizeof(user))	// 메모리 정렬에 의해 40 출력
}