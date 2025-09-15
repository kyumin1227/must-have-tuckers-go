package main

import "fmt"

func main() {
	var v1 int = 3
	var v2 interface{} = &v1
	var v3 int = *(v2.(*int))

	// 빈 인터페이스를 박싱 후 언박싱 하면 서로 다른 객체임
	fmt.Printf("v1: %x %T\n", &v1, &v1)
	fmt.Printf("v2: %x %T\n", &v2, &v2)
	fmt.Printf("v3: %x %T\n", &v3, &v3)
}
