package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a)	// 범위가 작은 타입으로 변환 시 상위 바이트가 사라짐 (현재의 경우 상위 1바이트)

	fmt.Println(a)
	fmt.Println(c)
}