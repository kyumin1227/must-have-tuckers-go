package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	// 8 바이트보다 작은 필드는 몰아서 배치할 경우 효율적
	A	int8	// 1
	C	int8	// 1
	E	int8	// 1
	B	int		// 8
	D	int		// 8
}

func main() {
	user := User{ 1, 2, 3, 4, 5 }
	fmt.Println(unsafe.Sizeof(user))	// 24 출력
}